package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type StorageCategory interface {
	CreateCategory(ctx context.Context, providerSign string, url string) (int, error)
	GetCategory(ctx context.Context) ([]string, error)
}

type CategoryRepo struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewCategoryRepo(db *pgxpool.Pool, log *zap.Logger) *CategoryRepo {
	return &CategoryRepo{
		db:     db,
		logger: log,
	}
}

func (r *CategoryRepo) CreateCategory(ctx context.Context, providerSign string, url string) (int, error) {
	return 0, nil
}

func (r *CategoryRepo) GetCategory(ctx context.Context) ([]string, error) {
	return nil, nil
}
