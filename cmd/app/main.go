package main

import (
	"fmt"

	"github.com/pooulad/makeFoD/pkg"
	"github.com/pooulad/makeFoD/util"
)

func main() {
	flag, err := pkg.ReadFlag()
	util.CheckError(err)

	for i := 1; i < count+1; i++ {
		dirName := fmt.Sprint(dir, i)
		os.Mkdir(dirName, 0777)
		// checkError(err)

		f, err := os.Create(fmt.Sprintf("./%v/%v.%v", dirName, fileName, extention))
		checkError(err)

		_, err = f.Write([]byte(fmt.Sprintf("# %v", dirName)))
		checkError(err)

		fmt.Println("new file created")
		defer f.Close()
	}
}
