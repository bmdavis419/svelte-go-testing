package storage

import (
	"github.com/jmoiron/sqlx"
)

type TodoStorage struct {
	Conn *sqlx.DB
}

type Todo_DB struct {
	Title       string
	Description string
	Id          int
	Completed   bool
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
	res, err := s.Conn.Exec(stmt, data.Title, data.Description, false, data.UserId)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *TodoStorage) GetAllTodos(userId int) ([]Todo_DB, error) {
	todos := []Todo_DB{}
	err := s.Conn.Select(&todos, "SELECT title, description, id, completed FROM todos WHERE user_id=?", userId)
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (s *TodoStorage) GetOneTodo(todoId int) (Todo_DB, error) {
	todo := Todo_DB{}
	err := s.Conn.Get(&todo, "SELECT title, description, id, completed FROM todos WHERE id=?", todoId)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (s *TodoStorage) CompleteTodo(todoId int) error {
	_, err := s.Conn.Exec("UPDATE todos SET completed=true WHERE id=?", todoId)
	return err
}

func (s *TodoStorage) DeleteTodo(todoId int) error {
	_, err := s.Conn.Exec("DELETE FROM todos WHERE id=?", todoId)
	return err
}
