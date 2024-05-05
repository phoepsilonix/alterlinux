package boot

import "errors"

type Mode int

var (
	BiosSyslinuxMbr             Mode = 0
	BiosSyslinuxElTorito        Mode = 1
	UefiIa32SystemdBootEsp      Mode = 2
	UefiX64SystemdBootEsp       Mode = 3
	UefiIa32SystemdBootElTorito Mode = 4
	UefiX64SystemdBootElTorito  Mode = 5
	UefiX64GrubEsp              Mode = 6
	UefiIa32GrubEsp             Mode = 7
	UefiX64GrubElTorito         Mode = 8
	UefiIa32GrubElTorito        Mode = 9
)

var ErrInvalidMode = errors.New("invalid boot mode")

func getModeFromStr(mode string) (Mode, error) {
	switch mode {
	case "bios.syslinux.mbr":
		return BiosSyslinuxMbr, nil
	case "bios.syslinux.eltorito":
		return BiosSyslinuxElTorito, nil
	case "uefi-ia32.systemd-boot.esp":
		return UefiIa32SystemdBootEsp, nil
	case "uefi-x64.systemd-boot.esp":
		return UefiX64SystemdBootEsp, nil
	case "uefi-ia32.systemd-boot.eltorito":
		return UefiIa32SystemdBootElTorito, nil
	case "uefi-x64.systemd-boot.eltorito":
		return UefiX64SystemdBootElTorito, nil
	case "uefi-x64.grub.esp":
		return UefiX64GrubEsp, nil
	case "uefi-ia32.grub.esp":
		return UefiIa32GrubEsp, nil
	case "uefi-x64.grub.eltorito":
		return UefiX64GrubElTorito, nil
	case "uefi-ia32.grub.eltorito":
		return UefiIa32GrubElTorito, nil

	}
	return 0, ErrInvalidMode
}

func GetModes(modes ...string) ([]Mode, error) {
	r := []Mode{}
	for _, mode := range modes {
		m, err := getModeFromStr(mode)
		if err != nil {
			return nil, err
		}
		r = append(r, m)
	}
	return r, nil
}
