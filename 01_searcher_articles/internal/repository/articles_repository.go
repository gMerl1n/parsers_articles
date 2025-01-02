package repository

import (
	"context"
	"fmt"

	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type ArticleRepo struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewArticleRepo(db *pgxpool.Pool, log *zap.Logger) *ArticleRepo {
	return &ArticleRepo{
		db:     db,
		logger: log,
	}
}

const (
	tableArticles = "articles"
)

type RepoArticle interface {
	CreateArticle(ctx context.Context, article entities.Article) (bool, error)
	CreateArticles(ctx context.Context, articles []entities.Article) (bool, error)
}

func (a *ArticleRepo) CreateArticle(ctx context.Context, article entities.Article) (bool, error) {

	var id int

	query := fmt.Sprintf("INSERT INTO %s (author, title, body, url, provider_sign, published_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", tableArticles)

	if err := a.db.QueryRow(ctx, query, article.Author, article.Title, article.Body, article.URL, article.ProviderSign, article.PublishedAt).Scan(&id); err != nil {
		a.logger.Warn("Failed to insert a new article into the table")
		return false, err
	}

	return true, nil

}

func (a *ArticleRepo) CreateArticles(ctx context.Context, articles []entities.Article) (bool, error) {

	for _, article := range articles {
		_, err := a.CreateArticle(ctx, article)
		if err != nil {
			a.logger.Warn("Failed to insert a new article into the table")
			fmt.Println(err)
		}

	}

	return true, nil

}
