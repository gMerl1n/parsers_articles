package domain

import "time"

type Article struct {
	ID           int
	Title        string
	Author       string
	Body         string
	ProviderSign string
	URL          string
	CategoryName string
	PublishedAt  time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
