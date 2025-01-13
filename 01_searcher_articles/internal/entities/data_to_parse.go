package entities

type DataForParsing struct {
	IDCategory  int
	UrlCategory string
	Provider    string
	Articles    []Article
	Deep        int64
}

func NewDataForParsing(urlCategory, provider string, deep int64, IDCategory int) *DataForParsing {

	sliceArticles := make([]Article, 0)

	return &DataForParsing{
		IDCategory:  IDCategory,
		UrlCategory: urlCategory,
		Provider:    provider,
		Articles:    sliceArticles,
		Deep:        deep,
	}
}
