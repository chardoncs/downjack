package licenses

import (
	"strings"

	"go.chardoncs.dev/downjack/internal/licenses/regex/ext"
)

type MatchedItem struct {
	Id       string
	Filename string
}

type SearchResult struct {
	Items []MatchedItem
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
		id := GetLicenseId(filename)
		loweredId := strings.ToLower(id)

		if lowerKeyword == loweredId {
			item := MatchedItem{
				Id:       id,
				Filename: filename,
			}

			result.Items = append(result.Items, item)
		}
	}

	return result, nil
}

func GetLicenseId(filename string) (result string) {
	result, _ = strings.CutSuffix(filename, ".tmpl")
	result = ext.GetRecognizedExtPattern().
		ReplaceAllString(result, "")

	return result
}
