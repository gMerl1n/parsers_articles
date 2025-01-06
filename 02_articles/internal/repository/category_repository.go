package repository

import (
	"context"
	"fmt"

	"github.com/gMerl1on/parsers_articles/02_articles/internal/domain"
	er "github.com/gMerl1on/parsers_articles/02_articles/pkg/errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const (
	categoryTable = "categories"
)

type StorageCategory interface {
	CreateCategory(ctx context.Context, name, url, providerSign string) (int, error)
	GetCategories(ctx context.Context) ([]domain.Category, error)
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

func (r *CategoryRepo) CreateCategory(ctx context.Context, name, url, providerSign string) (int, error) {

	var categoryID int

	query := fmt.Sprintf("INSERT INTO %s (name, url, provider_sign) VALUES ($1, $2, $3) RETURNING id", categoryTable)

	if err := r.db.QueryRow(ctx, query,
		name,
		url,
		providerSign,
	).Scan(&categoryID); err != nil {
		return 0, er.IncorrectRequest.SetCause(fmt.Sprint(err))
	}

	return categoryID, nil

}

func (r *CategoryRepo) GetCategories(ctx context.Context) ([]domain.Category, error) {

	categories := make([]domain.Category, 0)

	query := fmt.Sprintf("SELECT id, name, url, provider_sign, created_at FROM %s", categoryTable)

	rowsCategories, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rowsCategories.Next() {
		var category domain.Category

		err = rowsCategories.Scan(
			&category.ID,
			&category.Name,
			&category.ProviderSign,
			&category.URL,
			&category.CreatedAt,
		)

		if err != nil {
			er.IncorrectRequest.SetCause(fmt.Sprint(err))
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
