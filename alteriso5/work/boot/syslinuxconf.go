package boot

import (
	"os"
	"path"

	"github.com/Hayao0819/nahi/tputils"
)

type SyslinuxConf struct {
	Base string
}

func ReadSysLinuxConf(dir string) (*SyslinuxConf, error) {
	return &SyslinuxConf{
		Base: dir,
	}, nil
}

func (s *SyslinuxConf) ParseAndBuild(data any, out string) error {
	files, err := os.ReadDir(s.Base)
	if err != nil {
		return err
	}

	for _, file := range files {
		f := path.Join(s.Base, file.Name())
		buf, err := tputils.ApplyTemplate(f, data)
		if err != nil {
			return err
		}

		if err := os.WriteFile(path.Join(out, file.Name()), buf.Bytes(), 0644); err != nil {
			return err
		}
	}

	return nil
}
