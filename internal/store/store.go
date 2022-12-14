package store

import (
	"context"

	"github.com/go-redis/redis/v9"
)

// Store type manages the client/redis functions.
type Store struct {
	Conn *redis.Client
}

// Put
// adds a new pair to redis.
func (s *Store) Put(key, value string) error {
	ctx := context.Background()

	if err := s.Conn.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}

	return nil
}

// Trash
// removes a pair from redis.
func (s *Store) Trash(key string) error {
	ctx := context.Background()

	if err := s.Conn.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}

// Pull
// gets a pair from redis.
func (s *Store) Pull(key string) (string, error) {
	ctx := context.Background()

	value, err := s.Conn.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}
