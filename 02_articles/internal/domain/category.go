package domain

import "time"

type Category struct {
	ID           int
	Name         string
	ProviderSign string
	URL          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
