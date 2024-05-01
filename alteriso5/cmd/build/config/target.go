package config

type Target struct {
	Arch string
	Out  string
}

func NewTarget(arch string, out string) Target {
	return Target{arch, out}
}
