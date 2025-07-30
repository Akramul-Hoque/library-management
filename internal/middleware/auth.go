package middleware

import (
	"context"
	"library-management/internal/response"
	"net/http"
	"strings"
)

type contextKey string

var whitelist = map[string]bool{
	"/api/user/login": true,
	"/api/user":       true,
}

// AuthMiddleware checks JWT for all requests except whitelisted ones
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isWhitelisted(r) {
			next.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Universal(w, http.StatusUnauthorized, false, "Missing or invalid Authorization header", "NO_AUTH_HEADER", nil)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ParseJWT(tokenString)
		if err != nil {
			response.Universal(w, http.StatusUnauthorized, false, "Invalid or expired token", "INVALID_TOKEN", nil)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), contextKey("claims"), claims))
		next.ServeHTTP(w, r)
	})
}

func isWhitelisted(r *http.Request) bool {
	if whitelist[r.URL.Path] {
		if r.URL.Path == "/api/user" && r.Method != http.MethodPost {
			return false
		}
		return true
	}
	return false
}
