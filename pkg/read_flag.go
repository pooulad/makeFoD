package pkg

import (
	"errors"
	"flag"
)

type FlagReturns struct {
	Extention string
	DirName   string
	FileName  string
	DNumber   int
	FNumber   int
}

func ReadFlag() (*FlagReturns, error) {
	fr := FlagReturns{}
	flag.StringVar(&fr.Extention, "e", "", "get file format")
	flag.StringVar(&fr.FileName, "f", "", "get file name")
	flag.StringVar(&fr.DirName, "d", "", "get directory name")
	flag.IntVar(&fr.DNumber, "dn", 0, "get number of directories for create")
	flag.IntVar(&fr.FNumber, "fn", 0, "get number of files for create")
	flag.Parse()

	// if fr.FileName == "" {
	// 	if fr.DirName == "" {
	// 		return nil, errors.New("please enter file or directory name [-f or --f] or [-d or --d]")
	// 	}
	// } else {
	// 	if fr.FNumber == 0 {
	// 		return nil, errors.New("please enter file number to generate [-fn or --fn]")
	// 	}
	// }

	if fr.DirName == "" {
		if fr.FileName == "" {
			return nil, errors.New("please enter file or directory name [-f or --f] or [-d or --d]")
		} else {
			if fr.FNumber == 0 {
				return nil, errors.New("please enter file number to generate [-fn or --fn]")
			}
			if fr.Extention == "" {
				return nil, errors.New("please enter file format with tag [-e or --e]")
			}
		}
	} else {
		if fr.DNumber == 0 {
			return nil, errors.New("please enter directory number to generate [-dn or --dn]")
		}
		if fr.FileName != "" {
			if fr.FNumber == 0 {
				return nil, errors.New("please enter file number to generate [-fn or --fn]")
			}
			if fr.Extention == "" {
				return nil, errors.New("please enter file format with tag [-e or --e]")
			}
		}
	}

	return &fr, nil
}
