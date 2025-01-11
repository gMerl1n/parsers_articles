package domain

import "time"

type User struct {
	ID        int
	Name      string
	Surname   string
	Age       int
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
