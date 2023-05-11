package auth

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type SessionManager struct {
	Rdb  *redis.Client
	Conn *pgxpool.Pool
}

func NewSessionManager(rdb *redis.Client, conn *pgxpool.Pool) *SessionManager {
	return &SessionManager{Rdb: rdb, Conn: conn}
}

type UserSession struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (s *SessionManager) GenerateSession(data UserSession) (string, error) {
	sessionId := uuid.NewString()
	jsonData, _ := json.Marshal(data)
	err := s.Rdb.Set(context.Background(), sessionId, string(jsonData), 24*time.Hour).Err()
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

func (s *SessionManager) SignIn(email, password string) (string, error) {
	// check if the user exists
	var user User
	err := s.Conn.QueryRow(context.Background(), "select id, first_name, last_name, email, password from users where email = $1", email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return "", err
	}

	// check if the password matches
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// create the session
	sessionId := uuid.NewString()
	jsonData, _ := json.Marshal(UserSession{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	err = s.Rdb.Set(context.Background(), sessionId, string(jsonData), 24*time.Hour).Err()
	if err != nil {
		return "", err
	}

	return sessionId, nil
}

func (s *SessionManager) SignOut(sessionId string) error {
	return s.Rdb.Del(context.Background(), sessionId).Err()
}

func (s *SessionManager) GetSession(session string) (*UserSession, error) {
	data, err := s.Rdb.Get(context.Background(), session).Result()
	if err != nil {
		return nil, err
	}

	// unmarshal the data
	var userSession UserSession
	err = json.Unmarshal([]byte(data), &userSession)
	if err != nil {
		return nil, err
	}

	return &userSession, nil

}
