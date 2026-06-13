package api

import (
	"net/http"
	"strings"
)

const SecretToken = "chlorine12597-4587802ccd70540989555071f63d94a2b38fac719973d3e92ed8d56f7aee50d0424e0ca9d4dee6519d5c35035e56eb40dbb5528a98395d70fdb94e6099ea311d"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: Missing or malformed token", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != SecretToken {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}