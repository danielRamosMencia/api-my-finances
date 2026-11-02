package movementhistory

import (
	"encoding/json"
	"net/http"
)

func HldrGetUserHistory(w http.ResponseWriter, r *http.Request) {
	history, err := ServGetUserHistory(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(history)
}
