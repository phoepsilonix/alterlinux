package airootfs

func Get(name string) (FileSystem, error) {
	return SquashFS{}, nil
}
