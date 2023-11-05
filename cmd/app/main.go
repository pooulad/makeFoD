package main

import (
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	count := 4
	dir := "natas"
	fileName := "README"
	extention := "md"

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
