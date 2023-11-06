package main

import (
	"fmt"

	"github.com/pooulad/makeFoD/pkg"
	"github.com/pooulad/makeFoD/util"
)

func main() {
	flag, err := pkg.ReadFlag()
	util.FatalError(err)

	f := pkg.NewMakeFoD(flag)

	if flag.DirName == "" {
		err = f.CreateMultiFiles()
		if err != nil {
			util.FatalError(err)
		}
	} else {
		if flag.FileName == "" {
			err = f.CreateMultiEmptyDirectories()
			if err != nil {
				util.FatalError(err)
			}
		} else {
			err = f.CreateMultiDirectoriesWithFile()
			if err != nil {
				util.FatalError(err)
			}
		}
	}

	fmt.Println(flag)
	fmt.Println("ALL DONE SUCCESSFULY🦥")
}
