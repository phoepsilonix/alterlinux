package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type Profile struct {
	Base       string
	InstallDir string   `json:"install_dir"`
	BootModes  []string `json:"bootmodes"`
	ISOName    string   `json:"iso_name"`
	ISOLabel   string   `json:"iso_label"`
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

	fmt.Println(new)

	return &new, nil
}
