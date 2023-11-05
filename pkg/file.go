package pkg

import (
	"fmt"
	"os"
)

type MakeFoD struct {
	dirName  string
	fileName string
	fnumber  int
	dnumber  int
}

func NewMakeFoD(f *FlagReturns) *MakeFoD {
	return &MakeFoD{
		dirName:  f.DirName,
		dnumber:  f.DNumber,
		fnumber:  f.FNumber,
		fileName: f.FileName,
	}
}

