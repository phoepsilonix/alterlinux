package boot_test

import (
	"os"
	"path"
	"testing"

	"github.com/FascodeNet/alterlinux/alteriso5/work/boot"
)

func TestReadSysLinuxConf(t *testing.T) {
	wd, _ := os.Getwd()
	profile := path.Join(wd, "profile")
	conf, err := boot.ReadSysLinuxConf(profile)
	if err != nil {
		t.Fatal(err)
	}

	if conf.Base != profile {
		t.Errorf("expected %s, got %s", profile, conf.Base)
	}
}
