package utils

import "os"

func IsFileEmpty(fp *os.File) (bool, error) {
	stat, err := fp.Stat()
	if err != nil {
		return true, err
	}

	return stat.Size() == 0, nil
}
