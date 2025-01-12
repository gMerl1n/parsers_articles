package repository

import (
	"context"
	"fmt"

	"github.com/gMerl1on/parsers_articles/02_articles/constants"
	"github.com/gMerl1on/parsers_articles/02_articles/internal/domain"
	er "github.com/gMerl1on/parsers_articles/02_articles/pkg/errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	userTable = "users"
)

type StorageUser interface {
	CreateUser(ctx context.Context, name, surname, email, password string, age int) (int, error)
	GetUserByEmail(ctx context.Context, email, password string) (*domain.UserByEmail, error)
}

type UserRepo struct {
	db     *pgxpool.Pool
	redis  *redis.Client
	logger *zap.Logger
}

func NewUserRepo(db *pgxpool.Pool, r *redis.Client, log *zap.Logger) *UserRepo {
	return &UserRepo{
		db:     db,
		redis:  r,
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

func (u *UserRepo) GetUserByEmail(ctx context.Context, email, password string) (*domain.UserByEmail, error) {

	var userByEmail domain.UserByEmail

	query := fmt.Sprintf(`SELECT id, email, password FROM %s WHERE email = $1`, userTable)

	if err := u.db.QueryRow(ctx, query, email).Scan(&userByEmail.ID, &userByEmail.Email, &userByEmail.Password); err != nil {
		return &userByEmail, er.IncorrectRequest.SetCause(fmt.Sprint(err))
	}

	return &userByEmail, nil

}
