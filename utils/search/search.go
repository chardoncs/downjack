package search

import (
	"embed"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/chardoncs/downjack/utils/io"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type SearchResult struct {
	Filenames			[]string
	IsExact				bool
	// The content if it is an exact match
	ExactContent		string
}

func SearchEmbed(
	keyword string,
	root *embed.FS,
	dirPrefix string,
	suffix string,
) (*SearchResult, error) {
	result, err := exactMatchEmbedFile(keyword, root, dirPrefix, suffix)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return searchEmbedDir(keyword, root, dirPrefix)
		}

		return nil, err
	}

	return result, nil
}

func exactMatchEmbedFile(
	keyword string,
	root *embed.FS,
	dirPrefix string,
	suffix string,
) (*SearchResult, error) {
	filename, err := constructPlausibleFilename(keyword, suffix)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadEmbedToString(root, filepath.Join(dirPrefix, filename))
	if err != nil {
		return nil, err
	}

	return &SearchResult{
		Filenames: []string{ filename },
		IsExact: true,
		ExactContent: content,
	}, nil
}

func searchEmbedDir(keyword string, root *embed.FS, dirPrefix string) (*SearchResult, error) {
	dir, err := root.ReadDir(dirPrefix)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(dir))
	for i, name := range dir {
		names[i] = name.Name()
	}

	matched := searchWords(keyword, names)
	return &SearchResult{ Filenames: matched }, nil
}

func constructPlausibleFilename(name string, suffix string) (string, error) {
	caser := cases.Title(language.Und)
	suffix = strings.TrimSpace(suffix)
	chunks := strings.Split(name, ".")

	var sb strings.Builder

	for i, s := range chunks {
		isLast := i == len(chunks) - 1

		if isLast && suffix != "" && s == suffix {
			// Try writing suffix
			if _, err := sb.WriteString(s); err != nil {
				return "", err
			}
			goto EndProc
		}

		if _, err := sb.WriteString(caser.String(s)); err != nil {
			// Try writing capitalized
			if _, err := sb.WriteString(caser.String(s)); err != nil {
				return "", err
			}
		}

		if !isLast || suffix != "" && s != suffix {
			// Try writing "."
			if _, err := sb.WriteString("."); err != nil {
				return "", err
			}
		}
	}

	// Write suffix
	if suffix != "" {
		if _, err := sb.WriteString(suffix); err != nil {
			return "", err
		}
	}

EndProc:
	return sb.String(), nil
}

// (a dumb way to) search from an array of words
func searchWords(keyword string, arr []string) []string {
	lowered := strings.ToLower(keyword)
	result := make([]string, 0, len(arr))

	for _, name := range arr {
		if strings.Contains(strings.ToLower(name), lowered) {
			result = append(result, name)
		}
	}

	return result
}
