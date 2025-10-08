package licenses

type MatchedItem struct {
	Id			string
	Filename	string
}

type SearchResult struct {
	Items		[]MatchedItem
}

func SearchEmbed(keyword string) (*SearchResult, error) {
	return nil, nil
}
