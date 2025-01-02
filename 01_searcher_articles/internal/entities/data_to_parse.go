package entities

type DataForParsing struct {
	UrlCategory string
	Provider    string
	Articles    []Article
	Deep        int64
}

func NewDataForParsing(urlCategory, provider string, deep int64) *DataForParsing {

	sliceArticles := make([]Article, 0)

	return &DataForParsing{
		UrlCategory: urlCategory,
		Provider:    provider,
		Articles:    sliceArticles,
		Deep:        deep,
	}
}
