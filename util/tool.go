package util

import "log"

func FatalError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func PanicError(e error) {
	if e != nil {
		panic(e)
	}
}
