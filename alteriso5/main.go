package main

import (
	"fmt"
	"os"

	"github.com/FascodeNet/alterlinux/alteriso5/build"
)


func main(){
	err := build.Run()

	if err != nil{
		fmt.Fprintln(os.Stderr, err)
	}
}
