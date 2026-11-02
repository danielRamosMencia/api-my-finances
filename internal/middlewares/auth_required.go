package middlewares

import (
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

func AuthRequired(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Content-Type", "application/json")

	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
		return false
	}

	tokenString = tokenString[len("Bearer "):]

	err := utils.VerifyToken(tokenString)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return false
	}

	return true
}
