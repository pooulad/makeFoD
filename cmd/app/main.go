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
	count := 20
	dir := "temp"
	os.Mkdir(dir, 0777)
	for i := 0; i < count+1; i++ {
		os.Mkdir(fmt.Sprint(i), 0777)
		f, err := os.Create(fmt.Sprintf("./temp/file%v.txt", i))
		checkError(err)
		fmt.Println("hello go again...")
		fmt.Println(f)

		// defer f.Close()
	}
}
