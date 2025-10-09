package licenses

import "strings"

type MatchedItem struct {
	Id			string
	Filename	string
}

type SearchResult struct {
	Items		[]MatchedItem
	Exact		bool
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
				Id: getLicenseId(filename),
				Filename: filename,
			}

			if loweredName == lowerKeyword {
				result.Items = []MatchedItem{ item }
				result.Exact = true
				break
			}

			result.Items = append(result.Items, item)
		}
	}

	return result, nil
}

func getLicenseId(filename string) string {
	before, _, _ := strings.Cut(filename, ".")
	return before
}
