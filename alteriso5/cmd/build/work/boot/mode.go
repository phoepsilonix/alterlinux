package boot

type Mode struct {
	name         string
	install      func() error
	setupXorriso func()
}

func (m *Mode) SetInstall(f func() error) {
	m.install = f
}

var AllModes []*Mode = []*Mode{
	SysLinux,
}

func getModeByName(name string) *Mode {
	for _, m := range AllModes {
		if m.name == name {
			return m
		}
	}
	return nil
}
