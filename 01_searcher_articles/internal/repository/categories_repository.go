package repository

import (
	"context"
	"fmt"

	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/entities"
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

const (
	tableCatgories = "categories"
)

type RepoCategory interface {
	GetCategoriesBySign(ctx context.Context, sign string) ([]entities.Category, error)
}

func (c *CategoryRepo) GetCategoriesBySign(ctx context.Context, sign string) ([]entities.Category, error) {

	listCategories := make([]entities.Category, 0)

	query := fmt.Sprintf("SELECT * FROM %s WHERE provider_sign $2", tableCatgories)

	rowsCategoriesBySign, err := c.db.Query(ctx, query, sign)
	if err != nil {
		c.logger.Warn("Failed to get categories from DB", zap.String("Provider Sign %s", sign))
		return nil, err
	}

	for rowsCategoriesBySign.Next() {

		var category entities.Category

		err = rowsCategoriesBySign.Scan(
			&category.ID,
			&category.Name,
			&category.ProviderSign,
			&category.URL,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			c.logger.Warn("Failed to create list of Categories", zap.String("Provider Sign %s", sign))
			return nil, err
		}

		listCategories = append(listCategories, category)

	}

	return listCategories, nil
}
