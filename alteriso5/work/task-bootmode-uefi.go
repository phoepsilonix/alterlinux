package work

import "log/slog"

// Prepare configuration files for systemd-boot
var makeCommonSystemdBootConfig *BuildTask = NewBuildTask("makeCommonSystemdBootConfig", func(w Work) error {
	return nil
})

var makeCommonSystemdBoot *BuildTask = NewBuildTask("makeCommonSystemdBoot", func(w Work) error {
	return nil
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

	// Copy systemd-boot configuration files

	// shellx64.efi is picked up automatically when on /

	return nil
})

var makeUefiX64SystemdBootElTorito *BuildTask = NewBuildTask("makeUefiX64SystemdBootElTorito", func(w Work) error {
	return nil
})
