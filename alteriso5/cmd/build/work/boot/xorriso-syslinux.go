package boot

import (
	"path"
)

// SysLinux MBR El Torito
func (o *xorriso) SetArgsForSysLinuxElTorito() {
	arg := xorrisoArg{
		name: "SysLinuxEltorito",
	}

	arg.add("-eltorito-boot", "boot/syslinux/isolinux.bin")
	arg.add("-eltorito-catalog", "boot/syslinux/boot.cat")
	arg.add("-no-emul-boot", "-boot-load-size", "4", "-boot-info-table")

	o.addArg(&arg)
}

func (o *xorriso) SetArgsForSysLinuxMBRBios() {

	arg := xorrisoArg{
		name: "SysLinuxMBRBios",
	}

	arg.add("-isohybrid-mbr", path.Join(o.fsDir, "boot", "syslinux", "isohqpfx.bin"))
	arg.add("--mbr-force-bootable")
	arg.add("-partition_offset", "16")

	o.addArg(&arg)

}
