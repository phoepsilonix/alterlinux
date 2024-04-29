package airootfs

type FileSystem interface {
	Init() error
	GetPath() (string, error)
	Mount(path string) error
}

