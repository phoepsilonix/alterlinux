package conf

type Build struct {
	Arch               string
	OSName             string
	IsoName            string
	IsoLabel           string
	IsoPublisher       string
	IsoApplication     string
	IsoVersion         string
	InstallDir         string
	GpgKey             string
	BootSplash         bool
	ThemeName          string
	ChannelName        string
	TarBall            bool
	TarComp            string
	TarCompOpt         string
	SfsComp            string
	SfsCompOpt         string
	IncludeInfo        bool
	CustomizedSysLinux bool
	NorescureEntry     bool
	AurHelperCommand   string
	AurHelperPackage   string
	AurHelperDepends   []string
	AurHelperArgs      []string
	Kernel             string
	UserName           string
	Password           string
	UserShell          string
	Memtest86          bool
	CowSpace           string
	LocaleName         string
	NoChName           bool
	NoCheckVersion     bool
	NoEfi              bool
	NoISO              bool
	NoAur              bool
	NoPkgBuild         bool
	NoSigCheck         bool
	NoRmWork           bool

	GitVersion bool
}

type Debug struct {
	Debug       bool
	PacmanDebug bool
	Cleaning    bool
	NoConfirm   bool
	NoColor     bool
	WorkDir     string
	OutDir      string
	//Nodepend bool
}

type Config struct {
	Degug Debug
	Build Build
}
