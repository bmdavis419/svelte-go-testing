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
	_, err := s.Conn.Exec("insert into users (first_name, last_name, email, password) values (?, ?, ?, ?)", data.FirstName, data.LastName, data.Email, data.Password)
	if err != nil {
		return 0, err
	}
	var id int
	err = s.Conn.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
