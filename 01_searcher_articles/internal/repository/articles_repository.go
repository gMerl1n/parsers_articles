package repository

import (
	"context"

	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type ArticleRepo struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewArticleRepo(db *pgxpool.Pool, log *zap.Logger) *ArticleRepo {
	return &ArticleRepo{
		db:     db,
		logger: log,
	}
}

type RepoArticle interface {
	CreateArticles(ctx context.Context, articles []entities.Article) (bool, error)
}

func (a *ArticleRepo) CreateArticles(ctx context.Context, articles []entities.Article) (bool, error) {
	return true, nil
}
