package airootfs

type SquashFS struct {
	Path string
}

func (s SquashFS) GetPath() (string, error) {
	return s.Path, nil
}

func (s SquashFS) Init() error {
	return nil
}

func (s SquashFS) Mount(path string) error {
	return nil
}
