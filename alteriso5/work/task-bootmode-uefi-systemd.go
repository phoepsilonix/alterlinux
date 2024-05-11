package work

import (
	"log/slog"
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/utils"
	"github.com/FascodeNet/alterlinux/alteriso5/work/boot"
	"github.com/Hayao0819/nahi/cputils"
	"github.com/Hayao0819/nahi/exutils"
	osutils "github.com/Hayao0819/nahi/futils"
)

// Prepare configuration files for systemd-boot
var makeCommonSystemdBoot *BuildTask = NewBuildTask("makeCommonSystemdBoot", func(w Work) error {
	var total int64 = 0
	// TODO: Get ucode list

	// Calculate ESP size

	{
		efiboot_files := []string{}

		// For UEFI x64
		if w.profile.HasBootMode(boot.UefiX64SystemdBootEsp) || w.profile.HasBootMode(boot.UefiX64SystemdBootElTorito) {
			efiboot_files = append(efiboot_files,
				utils.Slash(w.Dirs.Pacstrap, "/usr/lib/systemd/boot/efi/systemd-bootx64.efi"),
				utils.Slash(w.Dirs.Pacstrap, "/usr/share/edk2-shell/x64/Shell_Full.efi"),
				utils.Slash(w.Dirs.Pacstrap, "/boot/memtest86+/memtest.efi"),
				utils.Slash(w.Dirs.Pacstrap, "/usr/share/licenses/spdx/GPL-2.0-only.txt"),
			)
		}

		// For UEFI ia32
		if w.profile.HasBootMode(boot.UefiIa32SystemdBootEsp) || w.profile.HasBootMode(boot.UefiIa32SystemdBootElTorito) {
			efiboot_files = append(efiboot_files,
				utils.Slash(w.Dirs.Pacstrap, "/usr/lib/systemd/boot/efi/systemd-bootia32.efi"),
				utils.Slash(w.Dirs.Pacstrap, "/usr/share/edk2-shell/ia32/Shell_Full.efi"),
			)
		}

		// For kernel
		{
			chroot, err := w.GetChroot()
			if err != nil {
				return err
			}
			kernels, err := chroot.FindKernels()
			if err != nil {
				return err
			}
			for _, k := range kernels {
				for _, f := range k.Files() {
					if osutils.Exists(f) {
						efiboot_files = append(efiboot_files, f)
					}
				}
			}
		}

		for _, f := range efiboot_files {
			if s, err := osutils.GetFileSize(f); err == nil {
				total += s
			} else {
				slog.Warn("Failed to get file size", "file", f, "error", err)
			}
		}

		slog.Debug("Found files for efiboot", "files", efiboot_files)

	}

	// For efiboot files
	if err := os.MkdirAll(w.Dirs.Efiboot, 0755); err != nil {
		return err
	}
	sizes, err := osutils.GetFileSizesInDir(w.Dirs.Efiboot)
	if err != nil {
		return err
	}
	for _, s := range sizes {
		total += s
	}

	slog.Debug("efiboot img size", "size", total)

	return boot.MakeEfiBootImg(path.Join(w.Base, "efiboot.img"), total)
})

var makeCommonSystemdBootConfig *BuildTask = NewBuildTask("makeCommonSystemdBootConfig", func(w Work) error {
	workLoaderDir := path.Join(w.Base, w.target.Arch, "loader")
	if err := os.MkdirAll(path.Join(workLoaderDir, "entries"), 0755); err != nil {
		return err
	}
	profileEfibootDir := path.Join(w.profile.Base, "efiboot")

	cpTasks := []cputils.CopyTask{
		{
			Source: path.Join(profileEfibootDir, "loader", "loader.conf"),
			Dest:   path.Join(workLoaderDir, "loader.conf"),
		},
		{
			Source: path.Join(profileEfibootDir, "loader", "entries"),
			Dest:   path.Join(workLoaderDir, "entries"),
		},
	}

	if err := cputils.CopyAll(cpTasks...); err != nil {
		return err
	}

	return nil
})

var makeCommonSystemdBootConfigIsofs *BuildTask = NewBuildTask("makeCommonSystemdBootConfigIsofs", func(w Work) error {
	return nil
})

var makeCommonSystemdBootConfigEsp *BuildTask = NewBuildTask("makeCommonSystemdBootConfigEsp", func(w Work) error {

	// mcopy -i "${efibootimg}" -s "${work_dir}/loader" ::/
	cmd := exutils.CommandWithStdio("mcopy", "-i", w.Files.EfibootImg, "-s", path.Join(w.Dirs.WorkArch, "loader"), "::/")
	return cmd.Run()
})

var makeUefiX64SystemdBootEsp *BuildTask = NewBuildTask("makeUefiX64SystemdBootEsp", func(w Work) error {
	slog.Info("Setting up systemd-boot for x64 UEFI booting...")

	// Prepare configuration files
	if err := w.RunOnce(makeCommonSystemdBootConfig); err != nil {
		return err
	}

	// Prepare a FAT image for the EFI system partition
	if err := w.RunOnce(makeCommonSystemdBoot); err != nil {
		return err
	}

	// Copy systemd-boot EFI binary to the default/fallback boot path
	{
		systemdBootEfi := utils.Slash(w.Dirs.Pacstrap, "/usr/lib/systemd/boot/efi/systemd-bootx64.efi")
		err := exutils.CommandWithStdio("mcopy", "-i", w.Files.EfibootImg, systemdBootEfi, "::/EFI/BOOT/BOOTx64.EFI").Run()
		if err != nil {
			return err
		}
	}

	// Copy systemd-boot configuration files
	if err := w.RunOnce(makeCommonSystemdBootConfigEsp); err != nil {
		return err
	}

	// shellx64.efi is picked up automatically when on /
	{
		shellEfi := utils.Slash(w.Dirs.Pacstrap, "/usr/share/edk2-shell/x64/Shell_Full.efi")
		err := exutils.CommandWithStdio("mcopy", "-i", w.Files.EfibootImg, shellEfi, "::/shellx64.efi").Run()
		if err != nil {
			return err
		}
	}

	return nil
})

var makeUefiX64SystemdBootElTorito *BuildTask = NewBuildTask("makeUefiX64SystemdBootElTorito", func(w Work) error {
	if err := w.RunOnce(makeCommonSystemdBootConfig); err != nil {
		return err
	}

	if err := w.RunOnce(makeUefiX64SystemdBootEsp); err != nil {
		return err
	}

	// Additionally set up systemd-boot in ISO 9660. This allows creating a medium for the live environment by using
	// manual partitioning and simply copying the ISO 9660 file system contents.
	// This is not related to El Torito booting and no firmware uses these files.
	slog.Info("Preparing an /EFI directory for the ISO 9660 file system...")
	if err := os.MkdirAll(path.Join(w.Dirs.Iso, "EFI", "BOOT"), 0755); err != nil {
		return err
	}

	// Copy systemd-boot EFI binary to the default/fallback boot path
	systemdBootEfi := utils.Slash(w.Dirs.Pacstrap, "/usr/lib/systemd/boot/efi/systemd-bootx64.efi")
	t := cputils.CopyTask{
		Source: systemdBootEfi,
		Dest:   path.Join(w.Dirs.Iso, "EFI", "BOOT", "BOOTx64.EFI"),
		Perm:   0644,
	}
	if err := cputils.CopyAll(t); err != nil {
		return err
	}

	if err := w.RunOnce(makeCommonSystemdBootConfigIsofs); err != nil {
		return err
	}

	return nil
})
