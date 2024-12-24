package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

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

type RepoCategory interface {
	GetCategoriesBySign(ctx context.Context, sign string, id_scheduler int) ([]string, error)
}

func (c *CategoryRepo) GetCategoriesBySign(ctx context.Context, sign string, id_scheduler int) ([]string, error) {
	return nil, nil
}
