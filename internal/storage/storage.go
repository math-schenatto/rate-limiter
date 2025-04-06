package storage

import "time"

type Storage interface {
	Increment(key string, expiration time.Duration) (int, error)
	Get(key string) (int, error)
	Block(key string, duration time.Duration) error
	IsBlocked(key string) (bool, error)
	Reset(key string) error
}
