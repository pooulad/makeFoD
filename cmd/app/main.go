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
	err = f.CreateMultiDirectoriesWithOneFile()
	if err != nil {
		panic(err)
	}
	fmt.Println(flag)
}
