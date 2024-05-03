package work

import "github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work/boot"

var makeOutFiles *BuildTask = NewBuildTask("makeOutFiles", func(w *Work) error {

	boot.Xorriso.SetArgsForSysLinuxElTorito()
	boot.Xorriso.SetArgsForSysLinuxElTorito()

	return boot.Xorriso.Build(w.GetDirs().Iso, w.target.Out)
})
