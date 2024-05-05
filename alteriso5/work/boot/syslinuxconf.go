package boot

import (
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/Hayao0819/nahi/fputils"
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

		// Determine if the file is a plain text file
		if file.IsDir() {
			continue
		}
		t, err := fputils.DetectFileType(f)
		if err != nil {
			return err
		}
		if !strings.HasPrefix(t, "text") {
			slog.Warn("Skipping non-text file", "file", f)
			continue
		}

		// Apply the template
		buf, err := tputils.ApplyTemplate(f, data)
		if err != nil {
			return err
		}

		// Write the file
		if err := os.WriteFile(path.Join(out, file.Name()), buf.Bytes(), 0644); err != nil {
			return err
		}
	}

	return nil
}
