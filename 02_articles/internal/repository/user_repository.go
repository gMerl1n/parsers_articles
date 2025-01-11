package repository

import (
	"context"
	"fmt"

	"github.com/gMerl1on/parsers_articles/02_articles/constants"
	er "github.com/gMerl1on/parsers_articles/02_articles/pkg/errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const (
	userTable = "users"
)

type StorageUser interface {
	CreateUser(ctx context.Context, name, surname, email, password string, age int) (int, error)
}

type UserRepo struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewUserRepo(db *pgxpool.Pool, log *zap.Logger) *UserRepo {
	return &UserRepo{
		db:     db,
		logger: log,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, name, surname, email, password string, age int) (int, error) {

	var userID int

	query := fmt.Sprintf(`INSERT INTO %s (name, surname, email, password, age, role_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, userTable)

	if err := u.db.QueryRow(ctx, query, name, surname, email, password, age, constants.UserRoleID).Scan(&userID); err != nil {
		return 0, er.IncorrectRequest.SetCause(fmt.Sprint(err))
	}

	return userID, nil

}
