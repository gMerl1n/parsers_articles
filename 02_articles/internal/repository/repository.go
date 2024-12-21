package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	Articles   StorageArticles
	Categories StorageCategory
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		Articles:   NewArticlesRepo(db),
		Categories: NewCategoryRepo(db),
	}
}
