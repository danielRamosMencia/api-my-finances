package movements

import (
	"encoding/json"
	"net/http"
)

func HldrGetUserMovements(w http.ResponseWriter, r *http.Request) {
	movements, err := ServGetUserMovements()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movements)
}
