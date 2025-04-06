package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort         string
	RateLimitIP        int
	RateLimitToken     int
	BlockDuration      time.Duration
	RedisAddr          string
	RedisPassword      string
	RedisDB            int
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	AppConfig = &Config{
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		RateLimitIP:    getEnvAsInt("RATE_LIMIT_IP", 10),
		RateLimitToken: getEnvAsInt("RATE_LIMIT_TOKEN", 100),
		BlockDuration:  time.Duration(getEnvAsInt("BLOCK_DURATION_SECONDS", 300)) * time.Second,
		RedisAddr:      getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:  getEnv("REDIS_PASSWORD", ""),
		RedisDB:        getEnvAsInt("REDIS_DB", 0),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valStr := getEnv(name, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}
