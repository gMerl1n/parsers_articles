package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repositories struct {
	Articles   StorageArticles
	Categories StorageCategory
}

func NewRepositories(db *pgxpool.Pool, log *zap.Logger) *Repositories {
	return &Repositories{
		Articles:   NewArticlesRepo(db, log),
		Categories: NewCategoryRepo(db, log),
	}
}
