package config

type Profile struct {
	Base string
}


func ReadProfile(path string) (Profile, error) {
	return Profile{}, nil
}

func DummyProfile() Profile {
	return Profile{}
}
