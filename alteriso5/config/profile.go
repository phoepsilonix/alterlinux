package config

import (
	"encoding/json"
	"os"
	"path"
	"slices"

	"github.com/FascodeNet/alterlinux/alteriso5/config/pkg"
	"github.com/FascodeNet/alterlinux/alteriso5/work/boot"
)

type Profile struct {
	Base             string
	InstallDir       string   `json:"install_dir"`
	BootModesStr     []string `json:"bootmodes"`
	BootModes        []*boot.Mode
	ISOName          string `json:"iso_name"`
	ISOLabel         string `json:"iso_label"`
	UseAlterSysLinux bool   `json:"use_alter_syslinux"`
}

func ReadProfile(dir string) (*Profile, error) {

	data, err := os.ReadFile(path.Join(dir, "profiledef.json"))
	if err != nil {
		return nil, err
	}

	new := Profile{
		Base: dir,
	}

	if err := json.Unmarshal(data, &new); err != nil {
		return nil, err
	}

	modes, err := boot.GetModes(new.BootModesStr...)
	if err != nil {
		return nil, err
	}
	new.BootModes = modes

	return &new, nil
}

func (p *Profile) GetPkgList(arch string) ([]string, error) {
	return pkg.GetPkgList(p.Base, arch)
}

func (p *Profile) HasBootMode(b *boot.Mode) bool {
	return slices.Contains(p.BootModesStr, b.String())
}
