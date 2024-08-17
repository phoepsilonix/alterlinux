package work

import (
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/utils"
	"github.com/Hayao0819/nahi/cputils"
)

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
		fmt.Printf("%v\n", w.Dirs.Pacstrap)
		systemdBootEfi := utils.Slash(w.Dirs.Pacstrap, "/usr/lib/systemd/boot/efi/systemd-bootx64.efi")
		err := utils.CommandWithStdio("mcopy", "-i", w.Files.EfibootImg, systemdBootEfi, "::/EFI/BOOT/BOOTx64.EFI").Run()
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
		err := utils.CommandWithStdio("mcopy", "-i", w.Files.EfibootImg, shellEfi, "::/shellx64.efi").Run()
		if err != nil {
			return err
		}
	}

	// Copy Memtest86+
	// TODO: Implement Memtest86+ support

	// Copy kernel and initramfs to FAT image.
	// systemd-boot can only access files from the EFI system partition it was launched from.
	//_run_once _make_boot_on_fat

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
