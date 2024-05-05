package boot

import "path"

var xorrisoCommonArgs = &xorrisoArg{
	args: func(o *xorriso) []string {
		out := path.Join(o.out, "alterlinux.iso")
		return []string{
			"-no_rc",
			"-as", "mkisofs",
			"-iso-level", "3",
			"-full-iso9660-filenames",
			"-joliet",
			"-joliet-long",
			"-rational-rock",
			"-volid", "ALTERLINUX",
			"--output", out,
			o.fsDir,
		}
	},
}
