package work

func New(dir string) (*Work, error) {
	return &Work{
		Base: dir,
	}, nil
}
