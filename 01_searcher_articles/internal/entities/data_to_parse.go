package entities

type DataForParsing struct {
	UrlCategory        string
	ProviderCategoryID string
	Provider           string
	Articles           []Article
	Deep               int
}
