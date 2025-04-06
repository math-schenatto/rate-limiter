package middleware

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/math-schenatto/rate-limiter/internal/limiter"
)

func RateLimitMiddleware(limiter *limiter.RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := getIP(r)
			token := r.Header.Get("API_KEY")

			log.Printf("Recebida requisição - IP: %s | Token: %s", ip, token)

			allowed, _, err := limiter.Check(ip, token)
			if err != nil {
				log.Printf("Erro ao verificar rate limit: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			if !allowed {
				log.Printf("Limite atingido para IP: %s ou Token: %s", ip, token)
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getIP(r *http.Request) string {
	// Se estiver atrás de proxy ou load balancer
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}

	// Padrão: r.RemoteAddr vem com IP:port, então usamos net.SplitHostPort
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
