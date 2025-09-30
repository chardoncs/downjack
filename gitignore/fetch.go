package gitignore

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const Repo = "github/gitignore"

type ReadInfo struct {
	Filename		string
	Content			string
}

func FetchEmbedded(name string) (*ReadInfo, error) {
	// Try exact match
	filename, err := constructPlausibleFilename(name)
	if err != nil {
		return nil, err
	}

	var file	fs.File

	file, err = root.Open(filepath.Join("./files", filename))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// Ensure file is nil
			file = nil
		} else {
			return nil, err
		}
	}

	// Try searching dir
	if file == nil {
		dir, err := root.ReadDir("./files")
		if err != nil {
			return nil, err
		}

		names := make([]string, len(dir))

		for i, entry := range dir {
			names[i] = entry.Name()
		}

		matched := searchWords(name, names)
		if matched == "" {
			return nil, os.ErrNotExist
		}

		file, err = root.Open()
		// TODO
	}

	// Process file
	defer file.Close()

	return &ReadInfo{}, nil
}

func FetchRepo(name string) (*ReadInfo, error) {
	// TODO
	return "", "", nil
}

func constructPlausibleFilename(name string) (string, error) {
	caser := cases.Title(language.Und)

	chunks := strings.Split(name, ".")

	var sb		strings.Builder

	for i, s := range chunks {
		isLast := i == len(chunks) - 1

		if isLast && s == "gitignore" {
			// Try writing "gitignore"
			if _, err := sb.WriteString(s); err != nil {
				return "", err
			}
			continue
		}

		if _, err := sb.WriteString(caser.String(s)); err != nil {
			// Try writing capitalized
			if _, err := sb.WriteString(caser.String(s)); err != nil {
				return "", err
			}
		}

		if !isLast {
			// Try writing "."
			if _, err := sb.WriteString("."); err != nil {
				return "", err
			}
		}
	}

	return sb.String(), nil
}

// (a dumb way to) search from an array of words
func searchWords(keyword string, arr []string) string {
	lowered := strings.ToLower(keyword)
}
