package config

type Profile struct {
	Base       string
	InstallDir string
}

func ReadProfile(path string) (Profile, error) {
	return Profile{}, nil
}
