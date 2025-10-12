package search 

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/chardoncs/downjack/utils"
	lib "github.com/chardoncs/downjack/internal/gitignore"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const gitignoreSuffix = ".gitignore"

type SearchResult struct {
	Filenames			[]string
	IsExact				bool
	// The content if it is an exact match
	ExactContent		string
}

func SearchEmbed(
	keyword string,
) (*SearchResult, error) {
	result, err := exactMatchEmbedFile(keyword)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return searchEmbedDir(keyword)
		}

		return nil, err
	}

	return result, nil
}

func exactMatchEmbedFile(
	keyword string,
) (*SearchResult, error) {
	filename, err := constructPlausibleFilename(keyword)
	if err != nil {
		return nil, err
	}

	content, err := utils.ReadEmbedToString(&lib.Root, filepath.Join(lib.DirPrefix, filename))
	if err != nil {
		return nil, err
	}

	return &SearchResult{
		Filenames: []string{ filename },
		IsExact: true,
		ExactContent: content,
	}, nil
}

func searchEmbedDir(keyword string) (*SearchResult, error) {
	dir, err := lib.Root.ReadDir(lib.DirPrefix)
	if err != nil {
		return nil, err
	}

	matched := searchWords(keyword, dir)
	return &SearchResult{ Filenames: matched }, nil
}

func constructPlausibleFilename(name string) (string, error) {
	caser := cases.Title(language.Und)
	name, _ = strings.CutSuffix(name, gitignoreSuffix)
	chunks := strings.Split(name, ".")

	var sb strings.Builder

	for _, s := range chunks {
		if _, err := sb.WriteString(caser.String(s)); err != nil {
			// Try writing capitalized
			if _, err := sb.WriteString(caser.String(s)); err != nil {
				return "", err
			}
		}

		if _, err := sb.WriteString("."); err != nil {
			return "", err
		}
	}

	if _, err := sb.WriteString("gitignore"); err != nil {
		return "", err
	}

	return sb.String(), nil
}

// (a dumb way to) search from an array of words
func searchWords(keyword string, dir []fs.DirEntry) []string {
	lowered := strings.ToLower(keyword)
	result := make([]string, 0, len(dir))

	for _, item := range dir {
		name := item.Name()

		if strings.Contains(strings.ToLower(name), lowered) {
			result = append(result, name)
		}
	}

	return result
}
