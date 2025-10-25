package utils

import (
	"embed"
	"os"
)

func IsFileEmpty(fp *os.File) (bool, error) {
	stat, err := fp.Stat()
	if err != nil {
		return true, err
	}

	return stat.Size() == 0, nil
}

func ListFilenames(root embed.FS, prefix string) ([]string, error) {
	dirEntries, err := root.ReadDir(prefix)
	if err != nil {
		return nil, err
	}

	strs := make([]string, len(dirEntries))
	for i, entry := range dirEntries {
		strs[i] = entry.Name()
	}

	return strs, nil
}
