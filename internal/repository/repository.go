package repository

import (
	"context"

	tt "github.com/gMerl1n/parsers_articles/internal/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	categoryTable = "categories"
)

type Storage interface {
	CreateCategory(ctx context.Context, providerSign string, url string) (int, error)
	GetCategory(ctx context.Context) ([]string, error)
	CreateArticle(ctx context.Context, article tt.Article) (string, error)
	GetArticles(ctx context.Context) ([]tt.Article, error)
}

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Storage {
	return &Repository{db: db}
}

func (r *Repository) CreateCategory(ctx context.Context, providerSign string, url string) (int, error) {
	return 0, nil
}

func (r *Repository) CreateArticle(ctx context.Context, article tt.Article) (string, error) {
	return "", nil
}

func (r *Repository) GetArticles(ctx context.Context) ([]tt.Article, error) {
	return nil, nil
}

func (r *Repository) GetCategory(ctx context.Context) ([]string, error) {
	return nil, nil
}
