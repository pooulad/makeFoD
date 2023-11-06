package pkg

import (
	"errors"
	"flag"
)

type FlagReturns struct {
	DirName  string
	FileName string
	DNumber  int
	FNumber  int
}

func ReadFlag() (*FlagReturns, error) {
	fr := FlagReturns{}
	flag.StringVar(&fr.FileName, "f", "", "get file name")
	flag.StringVar(&fr.DirName, "d", "", "get directory name")
	flag.IntVar(&fr.DNumber, "dn", 0, "get number of directories for create")
	flag.IntVar(&fr.FNumber, "fn", 0, "get number of files for create")
	flag.Parse()

	if fr.FileName == "" {
		if fr.DirName == "" {
			return nil, errors.New("please enter file or directory name")
		}
	} else {
		if fr.FNumber == 0 {
			return nil, errors.New("please enter file number to generate")
		}
	}

	if fr.DirName == "" {
		if fr.FileName == "" {
			return nil, errors.New("please enter file or directory name")
		}
	} else {
		if fr.DNumber == 0 {
			return nil, errors.New("please enter directory number to generate")
		}
	}

	return &fr, nil
}
