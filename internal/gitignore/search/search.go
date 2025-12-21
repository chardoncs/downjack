package search

import (
	"io/fs"
	"strings"

	lib "go.chardoncs.dev/downjack/internal/gitignore"
)

const gitignoreSuffix = ".gitignore"

type SearchResult struct {
	Filenames []string
}

func MatchFiles(keyword string) (*SearchResult, error) {
	dir, err := lib.Root.ReadDir(lib.DirPrefix)
	if err != nil {
		return nil, err
	}

	matched := searchWords(keyword, dir)
	return &SearchResult{Filenames: matched}, nil
}

func searchWords(keyword string, dir []fs.DirEntry) []string {
	lowered := strings.ToLower(keyword)
	result := make([]string, 0, len(dir))

	for _, item := range dir {
		name, _ := strings.CutSuffix(item.Name(), gitignoreSuffix)

		if strings.ToLower(name) == lowered {
			result = append(result, item.Name())
		}
	}

	return result
}
