package repository

import (
	avito "avito-backend-task-2023"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user avito.User) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (id, name, username, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, uuid.New().String(), user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserPostgres) DeleteUser(userId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	_, err := r.db.Exec(query, userId)

	return err
}
