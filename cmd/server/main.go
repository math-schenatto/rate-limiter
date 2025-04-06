package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/math-schenatto/rate-limiter/internal/config"
	"github.com/math-schenatto/rate-limiter/internal/limiter"
	"github.com/math-schenatto/rate-limiter/internal/middleware"
	"github.com/math-schenatto/rate-limiter/internal/storage"
)

func main() {
	config.LoadConfig()

	store := storage.NewRedisStorage(
		config.AppConfig.RedisAddr,
		config.AppConfig.RedisPassword,
		config.AppConfig.RedisDB,
	)

	rateLimiter := limiter.NewRateLimiter(store)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "✅ Hello! You're not rate-limited.")
	})
	handler := middleware.RateLimitMiddleware(rateLimiter)(mux)

	port := fmt.Sprintf(":%s", config.AppConfig.ServerPort)
	log.Printf("🟢 Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
