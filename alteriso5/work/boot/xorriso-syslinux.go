package boot

import (
	"path"
)

// SysLinux MBR El Torito
var xorrisoArgsForBiosSyslinuxElTorito = &xorrisoArg{
	bootMode: BiosSyslinuxElTorito,
	args: func(o *xorriso) []string {
		return []string{
			"-eltorito-boot", "boot/syslinux/isolinux.bin",
			"-eltorito-catalog", "boot/syslinux/boot.cat",
			"-no-emul-boot", "-boot-load-size", "4", "-boot-info-table",
		}
	},
}

var xorrisoArgsForBiosSyslinuxMbr = &xorrisoArg{
	bootMode: BiosSyslinuxMbr,
	args: func(o *xorriso) []string {
		return []string{
			"-isohybrid-mbr", path.Join(o.fsDir, "boot", "syslinux", "isohdpfx.bin"),
			"--mbr-force-bootable",
			"-partition_offset", "16",
		}
	},
}

func init() {
	Xorriso.addArgs(
		xorrisoArgsForBiosSyslinuxElTorito,
		xorrisoArgsForBiosSyslinuxMbr,
	)
}
