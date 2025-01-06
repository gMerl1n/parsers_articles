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
	articleTable = "articles"
)

type StorageArticles interface {
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

func (r *ArticleRepo) GetArticles(ctx context.Context) ([]domain.Article, error) {

	articles := make([]domain.Article, 0)

	query := fmt.Sprintf("SELECT id, author, title, body, url, provider_sign, published_at, created_at, updated_at  FROM %s", articleTable)

	rowsArticles, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rowsArticles.Next() {
		var article domain.Article

		err = rowsArticles.Scan(
			&article.ID,
			&article.Author,
			&article.Title,
			&article.Body,
			&article.URL,
			&article.ProviderSign,
			&article.PublishedAt,
			&article.CreatedAt,
			&article.UpdatedAt,
		)

		if err != nil {
			er.IncorrectRequest.SetCause(fmt.Sprint(err))
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil

}
