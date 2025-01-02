package entities

import "time"

type Article struct {
	Author       string
	Title        string
	ProviderSign string
	URL          string
	Body         string
	PublishedAt  time.Time
}

type Category struct {
	ID           int
	Name         string
	ProviderSign string
	URL          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
