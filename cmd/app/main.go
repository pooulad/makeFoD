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
	for i := 0; i < count; i++ {
		f, err := os.Create(fmt.Sprintf("file%v.txt", i))
		checkError(err)
		fmt.Println("hello go again...")
		fmt.Println(f)

		defer f.Close()
	}
}
