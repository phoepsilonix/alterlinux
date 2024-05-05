package config

import (
	"encoding/json"
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/config/pkg"
)

type Profile struct {
	Base             string
	InstallDir       string   `json:"install_dir"`
	BootModes        []string `json:"bootmodes"`
	ISOName          string   `json:"iso_name"`
	ISOLabel         string   `json:"iso_label"`
	UseAlterSysLinux bool     `json:"use_alter_syslinux"`
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

	return &new, nil
}

func (p *Profile) GetPkgList(arch string) ([]string, error) {
	return pkg.GetPkgList(p.Base, arch)
}
