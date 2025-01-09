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
	GetArticlesBySign(ctx context.Context, providerSign string) ([]domain.Article, error)
	GetArticlesByCategory(ctx context.Context, categoryID int) ([]domain.Article, error)
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

	query := "SELECT a.id, a.author, a.title, a.body, a.url, a.provider_sign, c.name, a.published_at, a.created_at, a.updated_at FROM articles AS a INNER JOIN categories AS c ON a.id = c.id"

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
			&article.CategoryName,
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

func (r *ArticleRepo) GetArticlesBySign(ctx context.Context, providerSign string) ([]domain.Article, error) {

	articlesBySign := make([]domain.Article, 0)

	query := fmt.Sprintf("SELECT id, author, title, body, url, provider_sign, published_at, created_at, updated_at  FROM %s WHERE provider_sign = $1", articleTable)

	rowsArticles, err := r.db.Query(ctx, query, providerSign)
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

		articlesBySign = append(articlesBySign, article)
	}

	return articlesBySign, nil

}

func (r *ArticleRepo) GetArticlesByCategory(ctx context.Context, categoryID int) ([]domain.Article, error) {

	var categoryArticles []domain.Article

	query := fmt.Sprintf(`SELECT a.id, a.author, a.title, a.body, a.url, c.name, a.provider_sign, a.published_at
			  			  FROM %s AS a JOIN %s AS c ON $1 = a.category_id
			  			  WHERE c.id = $1`, articleTable, categoryTable)

	rowsArticles, err := r.db.Query(ctx, query, categoryID)
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
			&article.CategoryName,
			&article.ProviderSign,
			&article.PublishedAt,
		)

		if err != nil {
			er.IncorrectRequest.SetCause(fmt.Sprint(err))
			return nil, err
		}

		categoryArticles = append(categoryArticles, article)

	}

	return categoryArticles, nil

}
