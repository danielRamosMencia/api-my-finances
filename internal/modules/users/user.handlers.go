package users

import (
	"encoding/json"
	"net/http"
)

func HdlrGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := ServGetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}
