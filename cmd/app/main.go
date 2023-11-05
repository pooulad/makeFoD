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
		fmt.Println(fmt.Sprint(dir, i))
		os.Mkdir(fmt.Sprint(dir, i), 0777)
		f, err := os.Create(fmt.Sprintf("./%v/%v.%v", fmt.Sprint(dir, i), fileName, extention))
		checkError(err)
		fmt.Println("new file created")
		fmt.Println(f)

		defer f.Close()
	}
}
