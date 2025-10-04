package gitignore

import "path/filepath"

func ReadToString(filename string) (string, error) {
	fullPath := filepath.Join(dirPrefix, filename)

	content, err := root.ReadFile(fullPath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
