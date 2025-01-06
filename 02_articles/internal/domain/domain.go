package domain

import "time"

type Article struct {
	ID           int
	Title        string
	Author       string
	Body         string
	ProviderSign string
	URL          string
	PublishedAt  time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Category struct {
	ID           int
	Name         string
	ProviderSign string
	URL          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
