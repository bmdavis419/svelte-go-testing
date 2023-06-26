package storage

import (
	"github.com/jmoiron/sqlx"
)

type UserStorage struct {
	Conn *sqlx.DB
}

func NewUserStorage(conn *sqlx.DB) *UserStorage {
	return &UserStorage{Conn: conn}
}

type NewUser struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (s *UserStorage) CreateNewUser(data NewUser) (int, error) {
	res, err := s.Conn.Exec("insert into users (first_name, last_name, email, password) values (?, ?, ?, ?)", data.FirstName, data.LastName, data.Email, data.Password)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
