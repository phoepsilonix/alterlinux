package boot

var SysLinux *Mode = &Mode{
	name: "SysLinux",
	setupXorriso: func() {
		Xorriso.setArgsForSysLinuxElTorito()
		Xorriso.setArgsForSysLinuxMBRBios()
	},
}
