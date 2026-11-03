package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
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

			if pattern == "" {
				responses.Error(w, http.StatusNotFound, responses.ROUTE_NOT_FOUND)
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
			responses.Error(w, http.StatusUnauthorized, responses.MISSING_TOKEN)
			return
		}

		tokenString, found := strings.CutPrefix(header, "Bearer ")
		if !found || tokenString == "" {
			responses.Error(w, http.StatusUnauthorized, responses.MISSING_TOKEN)
			return
		}

		claims, err := utils.VerifyToken(tokenString)
		if errors.Is(err, jwt.ErrTokenExpired) {
			responses.Error(w, http.StatusUnauthorized, responses.TOKEN_EXPIRED)
			return
		}

		if err != nil {
			responses.Error(w, http.StatusUnauthorized, responses.INVALID_TOKEN)
			return
		}

		session, _ := claims["session"].(string)
		if !sessions.IsActive(session, tokenString) {
			responses.Error(w, http.StatusUnauthorized, responses.SESSION_CLOSED)
			return
		}

		ctx := utils.ContextWithClaims(r.Context(), claims)
		mux.ServeHTTP(w, r.WithContext(ctx))
	})
}
