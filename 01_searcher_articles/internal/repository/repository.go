package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repositories struct {
	Categories RepoCategory
	Articles   RepoArticle
}

func NewRepositories(db *pgxpool.Pool, log *zap.Logger) *Repositories {
	return &Repositories{
		Categories: NewCategoryRepo(db, log),
		Articles:   NewArticleRepo(db, log),
	}
}
