package repository

import (
	"context"

	"github.com/gMerl1on/parsers_articles/02_articles/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type StorageArticles interface {
	CreateArticle(ctx context.Context, article domain.Article) (string, error)
	GetArticles(ctx context.Context) ([]domain.Article, error)
}

type ArticleRepo struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewArticlesRepo(db *pgxpool.Pool, log *zap.Logger) *ArticleRepo {
	return &ArticleRepo{
		db:     db,
		logger: log,
	}
}

func (r *ArticleRepo) CreateArticle(ctx context.Context, article domain.Article) (string, error) {
	return "", nil
}

func (r *ArticleRepo) GetArticles(ctx context.Context) ([]domain.Article, error) {
	return nil, nil
}
