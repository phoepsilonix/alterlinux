package utils

import (
	"io/fs"
	"os"
	"path/filepath"

	cp "github.com/otiai10/copy"
)

type CopyTask struct {
	Source string
	Dest   string
	Perm   fs.FileMode
	Skip   func(srcinfo os.FileInfo, src, dest string) (bool, error) // Skipするファイルならtrueを返す
}

func (c *CopyTask) Copy() error {

	opt := cp.Options{}

	if c.Perm != 0 {
		opt.PermissionControl = cp.AddPermission(c.Perm)
	}

	return cp.Copy(c.Source, c.Dest, opt)
}

func OnlySpecificExtention(ext string) func(srcinfo os.FileInfo, src, dest string) (bool, error) {
	return func(srcinfo os.FileInfo, src, dest string) (bool, error) {
		if srcinfo.IsDir() {
			return false, nil
		}
		if filepath.Ext(src) != ext {
			return true, nil
		}
		return false, nil
	}
}

func CopyAll(tasks ...CopyTask) error {
	for _, task := range tasks {
		if err := task.Copy(); err != nil {
			return err
		}
	}
	return nil
}
