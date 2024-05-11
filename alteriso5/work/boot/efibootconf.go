package boot

type EfibootConf struct {
	Base string
}

func ReadEfibootConf(dir string) (*EfibootConf, error) {
	return &EfibootConf{
		Base: dir,
	}, nil
}

func (e *EfibootConf) ParseAndBuild(data any, out string) error {
	return nil
}
