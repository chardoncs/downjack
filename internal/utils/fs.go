package utils

import (
	"embed"
	"errors"
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

func TryExactMatchFile(root embed.FS, fpath string) ([]byte, bool, error) {
	b, err := root.ReadFile(fpath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, false, nil
		}

		return nil, false, err
	}

	return b, true, nil
}
