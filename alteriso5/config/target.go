package config

type Target struct {
	Arch string
	Out  string
	// ProfileDir string
}

func NewTarget(arch string, out string) Target {
	return Target{
		Arch: arch,
		Out: out,
		// ProfileDir: profile,
	}
}
