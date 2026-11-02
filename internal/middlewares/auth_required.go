package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/danielRamosMencia/api-my-finances/internal/sessions"
	"github.com/danielRamosMencia/api-my-finances/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(mux *http.ServeMux, publicRoutes ...string) http.Handler {
	public := make(map[string]bool, len(publicRoutes))
	for _, route := range publicRoutes {
		public[route] = true
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, pattern := mux.Handler(r)

		if pattern == "" || public[pattern] {
			if !RequestLimiter(w, r, false) {
				return
			}

			mux.ServeHTTP(w, r)
			return
		}

		if !RequestLimiter(w, r, true) {
			return
		}

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString, found := strings.CutPrefix(header, "Bearer ")
		if !found || tokenString == "" {
			http.Error(w, "authorization header must be: Bearer <token>", http.StatusUnauthorized)
			return
		}

		claims, err := utils.VerifyToken(tokenString)
		if errors.Is(err, jwt.ErrTokenExpired) {
			http.Error(w, "token expired, please log in again", http.StatusUnauthorized)
			return
		}

		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		session, _ := claims["session"].(string)
		if !sessions.IsActive(session, tokenString) {
			http.Error(w, "session closed or not found, please log in again", http.StatusUnauthorized)
			return
		}

		ctx := utils.ContextWithClaims(r.Context(), claims)
		mux.ServeHTTP(w, r.WithContext(ctx))
	})
}
