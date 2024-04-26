package build

import (
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/config"
	"github.com/FascodeNet/alterlinux/alteriso5/work"
)


func Run()error{

	current, err := os.Getwd()

	if err != nil{
		return err
	}

	work, err := work.New(path.Join(current, "work"))
	if err != nil{
		return err
	}


	p := config.DummyProfile()
	return work.Build(p)
}
