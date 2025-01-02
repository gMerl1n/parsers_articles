package domain

type Article struct {
	Title   string
	Author  string
	Body    string
	Created float64
}

type Category struct {
	ID           int
	Name         string
	ProviderSign string
	URL          string
	CreatedAt    float64
	UpdatedAt    float64
}
