package entities

type Article struct {
	Author       string
	Title        string
	ProviderSign string
	URL          string
	Body         string
	PublishedAt  int64
}

type Category struct {
	ID           int
	Name         string
	ProviderSign string
	URL          string
	CreatedAt    float64
	UpdatedAt    float64
}
