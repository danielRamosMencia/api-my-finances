package auth

import (
	"encoding/json"
	"net/http"
)

func HldrLogin(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jwt, err := ServLogin(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(jwt)
}

func HldrLogout(w http.ResponseWriter, r *http.Request) {
	ServLogout(r.Context())

	w.WriteHeader(http.StatusNoContent)
}
