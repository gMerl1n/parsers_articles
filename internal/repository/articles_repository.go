package repository

import (
	"context"

	tt "github.com/gMerl1n/parsers_articles/internal/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageArticles interface {
	CreateArticle(ctx context.Context, article tt.Article) (string, error)
	GetArticles(ctx context.Context) ([]tt.Article, error)
}

type ArticleRepo struct {
	db *pgxpool.Pool
}

func NewArticlesRepo(db *pgxpool.Pool) *ArticleRepo {
	return &ArticleRepo{
		db: db,
	}
}

func (r *ArticleRepo) CreateArticle(ctx context.Context, article tt.Article) (string, error) {
	return "", nil
}

func (r *ArticleRepo) GetArticles(ctx context.Context) ([]tt.Article, error) {
	return nil, nil
}
