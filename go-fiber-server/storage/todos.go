package storage

import "github.com/jmoiron/sqlx"

type TodoStorage struct {
	Conn *sqlx.DB
}

func NewTodoStorage(conn *sqlx.DB) *TodoStorage {
	return &TodoStorage{Conn: conn}
}

type NewTodoInput struct {
	Title       string
	Description string
	UserId      int
}

func (s *TodoStorage) CreateNewTodo(data NewTodoInput) (int, error) {
	stmt := "INSERT INTO todos (title, description, completed, user_id) VALUES (?, ?, ?, ?)"
	_, err := s.Conn.Exec(stmt, data.Title, data.Description, false, data.UserId)
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
