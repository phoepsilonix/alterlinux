package boot

func GetModesByName(targets ...string) []*Mode {
	modes := []*Mode{}
	for _, t := range targets {
		m := getModeByName(t)
		if m != nil {
			modes = append(modes, m)
		}
	}
	return modes
}

func Setup(modes *[]*Mode) error {
	for _, m := range *modes {
		m.setupXorriso()
		if err := m.install(); err != nil {
			return err
		}
	}
	return nil
}
