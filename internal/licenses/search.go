package licenses

import (
	"strings"

	"github.com/chardoncs/downjack/internal/licenses/regex/ext"
)

type MatchedItem struct {
	Id       string
	Filename string
}

type SearchResult struct {
	Items []MatchedItem
	Exact bool
}

func SearchEmbed(keyword string) (*SearchResult, error) {
	dir, err := Root.ReadDir(DirPrefix)
	if err != nil {
		return nil, err
	}

	lowerKeyword := strings.ToLower(keyword)
	result := &SearchResult{
		Items: make([]MatchedItem, 0, len(dir)),
	}

	for _, entry := range dir {
		filename := entry.Name()
		loweredName := strings.ToLower(filename)

		if strings.Contains(loweredName, lowerKeyword) {
			item := MatchedItem{
				Id:       getLicenseId(filename),
				Filename: filename,
			}

			if strings.ToLower(item.Id) == lowerKeyword {
				result.Items = []MatchedItem{item}
				result.Exact = true
				break
			}

			result.Items = append(result.Items, item)
		}
	}

	return result, nil
}

func getLicenseId(filename string) (result string) {
	result, _ = strings.CutSuffix(filename, ".tmpl")
	result = ext.GetRecognizedExtPattern().
		ReplaceAllString(result, "")

	return result
}
