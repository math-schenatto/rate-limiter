package limiter

import (
	"fmt"
	"log"
	"time"

	"github.com/math-schenatto/rate-limiter/internal/config"
	"github.com/math-schenatto/rate-limiter/internal/storage"
)

type RateLimiter struct {
	storage storage.Storage
}

func NewRateLimiter(storage storage.Storage) *RateLimiter {
	return &RateLimiter{storage: storage}
}

func (r *RateLimiter) Allow(key string, limit LimitConfig) (bool, error) {
	blocked, err := r.storage.IsBlocked(key)
	if err != nil {
		return false, err
	}
	if blocked {
		return false, nil
	}

	count, err := r.storage.Increment(key, time.Second)
	if err != nil {
		return false, err
	}

	if count > limit.RequestsPerSecond {
		err := r.storage.Block(key, time.Duration(limit.BlockDurationSec)*time.Second)
		if err != nil {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

func (r *RateLimiter) Check(ip string, token string) (bool, string, error) {
	var key string
	var limit LimitConfig

	if token != "" {
		key = fmt.Sprintf("token:%s", token)
		limit = LimitConfig{
			RequestsPerSecond: config.AppConfig.RateLimitToken,
			BlockDurationSec:  int(config.AppConfig.BlockDuration.Seconds()),
		}
	} else {
		key = fmt.Sprintf("ip:%s", ip)
		limit = LimitConfig{
			RequestsPerSecond: config.AppConfig.RateLimitIP,
			BlockDurationSec:  int(config.AppConfig.BlockDuration.Seconds()),
		}
	}

	allowed, err := r.Allow(key, limit)
	if err != nil {
		log.Printf("❌ Error checking rate limit: %v", err)
	}
	return allowed, key, err
}
