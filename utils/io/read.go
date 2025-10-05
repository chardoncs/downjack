package io

import "embed"

func ReadEmbedToString(fs *embed.FS, fpath string) (string, error) {
	content, err := fs.ReadFile(fpath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
