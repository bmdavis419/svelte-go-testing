package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage struct {
	Conn *pgxpool.Pool
}

func NewUserStorage(conn *pgxpool.Pool) *UserStorage {
	return &UserStorage{Conn: conn}
}

func (s *UserStorage) GenerateNewUser() (int, error) {
	var id int
	err := s.Conn.QueryRow(context.Background(), "insert into users (first_name, last_name, email, password) values ($1, $2, $3, $4) returning id", "Test", "User", "test@gmail.com", "password").Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
