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

func (m *MakeFoD) CreateMultiDirectoriesWithOneFile() error {
	for i := 1; i < m.dnumber+1; i++ {
		dirName := fmt.Sprint(m.dirName, i)
		os.Mkdir(dirName, 0777)

		f, err := os.Create(fmt.Sprintf("./%v/%v", dirName, m.fileName))
		if err != nil {
			return err
		}

		defer f.Close()
	}

	return nil
}
