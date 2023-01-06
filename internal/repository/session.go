package repository

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type SessionRepository struct {
	client *redis.Client
}

func NewSessionRepository(client *redis.Client) *SessionRepository {
	return &SessionRepository{client: client}
}

func (s *SessionRepository) CreateSession(value, token string, exirationTime time.Duration) error {
	status := s.client.Set(token, interface{}(value), exirationTime)
	if status.Err() == redis.Nil {
		return fmt.Errorf("error when creating session: %w", status.Err())
	} else if status.Err() != nil {
		return fmt.Errorf("error when creating session: %w", status.Err())
	}

	return nil
}

func (s *SessionRepository) GetSession(token string) (string, error) {
	email := s.client.Get(token)

	if email.Err() == redis.Nil {
		return "", fmt.Errorf("token doesnt exists")
	} else if email.Err() != nil {
		return "", email.Err()
	}

	return email.String(), nil
}

func (s *SessionRepository) DeleteSession(token string) error {
	return s.client.Del(token).Err()
}
