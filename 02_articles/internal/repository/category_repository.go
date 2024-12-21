package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageCategory interface {
	CreateCategory(ctx context.Context, providerSign string, url string) (int, error)
	GetCategory(ctx context.Context) ([]string, error)
}

type CategoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) CreateCategory(ctx context.Context, providerSign string, url string) (int, error) {
	return 0, nil
}

func (r *CategoryRepo) GetCategory(ctx context.Context) ([]string, error) {
	return nil, nil
}
