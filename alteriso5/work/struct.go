package work

import "github.com/FascodeNet/alterlinux/alteriso5/config"

type Work struct {
	Base string
	Chroots []*Chroot
}





func (work *Work) InitChroot(arch string, cmd ...string)error{
	return nil;
}


func (work *Work) Build(p config.Profile)error{
	return nil;
}
