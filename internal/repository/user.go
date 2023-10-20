package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jatin510/go-chat-app/internal/models"
)

type UserRepoInterface interface {
	Create(models.User) (models.User, error)
	FindOneByEmail(string) (models.User, error)
	Update(models.User) (models.User, error)
	Delete(uuid.UUID) error
	FindOne(any) (models.User, error)
	FindAll(any) ([]models.User, error)
	Count(any) (int, error)
}

type user struct {
	db *pgx.Conn
	l  models.Logger
}

func NewUserRepo(db *pgx.Conn, l models.Logger) UserRepoInterface {
	return &user{
		db: db,
		l:  l,
	}
}

func (u user) Create(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBQueryTimeout)
	defer cancel()
	rows, err := u.db.Query(ctx, "INSERT INTO users VALUES ($1,$2,$3,$4,$5,$6)", user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		u.l.Error("error in Create room query", err.Error())
		return models.User{}, err
	}
	defer rows.Close()

	return user, nil
}

func (u user) FindOneByEmail(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBQueryTimeout)
	defer cancel()
	rows, err := u.db.Query(ctx, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		u.l.Error("error in Create room query", err.Error())
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (u user) Update(user models.User) (models.User, error) {
	return models.User{}, nil
}

func (u user) Delete(id uuid.UUID) error {
	return nil
}

func (u user) FindOne(filter any) (models.User, error) {
	return models.User{}, nil
}

func (u user) FindAll(filter any) ([]models.User, error) {
	return []models.User{}, nil
}

func (u user) Count(filter any) (int, error) {
	return 0, nil
}
