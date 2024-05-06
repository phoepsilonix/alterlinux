package boot

import "errors"

type Mode struct {
	name     string
	validate func() error
}

var (
	BiosSyslinuxMbr             *Mode = &Mode{"bios.syslinux.mbr", nil}
	BiosSyslinuxElTorito        *Mode = &Mode{"bios.syslinux.eltorito", nil}
	UefiIa32SystemdBootEsp      *Mode = &Mode{"uefi-ia32.systemd-boot.esp", nil}
	UefiX64SystemdBootEsp       *Mode = &Mode{"uefi-x64.systemd-boot.esp", nil}
	UefiIa32SystemdBootElTorito *Mode = &Mode{"uefi-ia32.systemd-boot.eltorito", nil}
	UefiX64SystemdBootElTorito  *Mode = &Mode{"uefi-x64.systemd-boot.eltorito", nil}
	UefiX64GrubEsp              *Mode = &Mode{"uefi-x64.grub.esp", nil}
	UefiIa32GrubEsp             *Mode = &Mode{"uefi-ia32.grub.esp", nil}
	UefiX64GrubElTorito         *Mode = &Mode{"uefi-x64.grub.eltorito", nil}
	UefiIa32GrubElTorito        *Mode = &Mode{"uefi-ia32.grub.eltorito", nil}

	Modes = []*Mode{
		BiosSyslinuxMbr,
		BiosSyslinuxElTorito,
		UefiIa32SystemdBootEsp,
		UefiX64SystemdBootEsp,
		UefiIa32SystemdBootElTorito,
		UefiX64SystemdBootElTorito,
		UefiX64GrubEsp,
		UefiIa32GrubEsp,
		UefiX64GrubElTorito,
		UefiIa32GrubElTorito,
	}
)

func (m *Mode) String() string {
	return m.name
}

func (m *Mode) Validate() error {
	if m.validate != nil {
		return m.validate()
	}
	return nil
}

var ErrInvalidMode = errors.New("invalid boot mode")

func getModeFromStr(mode string) (*Mode, error) {
	for _, m := range Modes {
		if m.name == mode {
			return m, nil
		}
	}
	return nil, ErrInvalidMode
}

func GetModes(modes ...string) ([]*Mode, error) {
	r := []*Mode{}
	for _, mode := range modes {
		m, err := getModeFromStr(mode)
		if err != nil {
			return nil, err
		}
		r = append(r, m)
	}
	return r, nil
}
