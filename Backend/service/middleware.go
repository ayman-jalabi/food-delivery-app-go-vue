package service

import (
	"net/http"
	"strings"
)

// ProtectedAPIRouteMiddleware will help check if client has a valid access token
func ProtectedAPIRouteMiddleware(tokenService *TokenService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		accessToken := parts[1]
		_, err := tokenService.ValidateAccessToken(accessToken)
		if err != nil {
			http.Error(w, "Invalid access token", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to the next handler
		next(w, r)
	}
}
