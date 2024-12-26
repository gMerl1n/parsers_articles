package entities

type DataForParsing struct {
	UrlCategory        string
	ProviderCategoryID string
	Provider           string
	Articles           []Article
	Deep               int64
}

func NewDataForParsing(urlCategory, providerCategoryID, provider string, deep int64) *DataForParsing {

	sliceArticles := make([]Article, 0)

	return &DataForParsing{
		UrlCategory:        urlCategory,
		ProviderCategoryID: provider,
		Provider:           provider,
		Articles:           sliceArticles,
		Deep:               deep,
	}
}
