// ratelimit.go

package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

// Limiter é uma instância do rate.Limiter
var Limiter = rate.NewLimiter(1, 5) // 1 requisição por segundo com um burst de 5

// RateLimitMiddleware é o middleware de limitação de taxa
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !Limiter.Allow() {
			http.Error(w, "Limite de taxa excedido", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
