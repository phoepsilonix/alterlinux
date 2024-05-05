package work

import "github.com/FascodeNet/alterlinux/alteriso5/work/boot"

// Run each bootmodes
var makeBootModes *BuildTask = NewBuildTask("makeBootModes", func(w Work) error {
	for _, mode := range w.profile.BootModes {
		switch mode.String() {
		case boot.BiosSyslinuxMbr.String():
			if err := w.RunOnce(makeBiosSysLinuxMbr); err != nil {
				return err
			}
		case boot.BiosSyslinuxElTorito.String():
			if err := w.RunOnce(makeBiosSysLinuxElTorito); err != nil {
				return err
			}
		case boot.UefiX64SystemdBootEsp.String():
			if err := w.RunOnce(makeUefiX64SystemdBootEsp); err != nil {
				return err
			}
		case boot.UefiX64SystemdBootElTorito.String():
			if err := w.RunOnce(makeUefiX64SystemdBootElTorito); err != nil {
				return err
			}
		}
	}

	return nil
})
