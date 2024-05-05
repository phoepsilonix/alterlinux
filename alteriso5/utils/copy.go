package utils

import (
	"io/fs"
	"log/slog"
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

	opt := cp.Options{
		Skip: c.Skip,
	}

	if c.Perm != 0 {
		opt.PermissionControl = cp.AddPermission(c.Perm)
	}

	return cp.Copy(c.Source, c.Dest, opt)
}

func OnlySpecificExtention(ext string) func(srcinfo os.FileInfo, src, dest string) (bool, error) {
	return func(srcinfo os.FileInfo, src, dest string) (bool, error) {
		//slog.Debug("Checking file", "file", src)
		if srcinfo.IsDir() {
			//slog.Debug("Skipping directory", "dir", src)
			return false, nil
		}
		if filepath.Ext(src) != ext {
			slog.Debug("Skipping file", "file", src, "ext", filepath.Ext(src))
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
