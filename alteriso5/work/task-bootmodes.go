package work

import "github.com/FascodeNet/alterlinux/alteriso5/work/boot"

// Run each bootmodes
var makeBootModes *BuildTask = NewBuildTask("makeBootModes", func(w Work) error {
	for _, mode := range w.profile.BootModes {
		switch mode {
		case boot.BiosSyslinuxMbr:
			if err := w.RunOnce(makeBiosSysLinuxMbr); err != nil {
				return err
			}
		case boot.BiosSyslinuxElTorito:
			if err := w.RunOnce(makeBiosSysLinuxElTorito); err != nil {
				return err
			}
		case boot.UefiX64SystemdBootEsp:
			if err := w.RunOnce(makeUefiX64SystemdBootEsp); err != nil {
				return err
			}
		case boot.UefiX64SystemdBootElTorito:
			if err := w.RunOnce(makeUefiX64SystemdBootElTorito); err != nil {
				return err
			}
		}
	}

	return nil
})
