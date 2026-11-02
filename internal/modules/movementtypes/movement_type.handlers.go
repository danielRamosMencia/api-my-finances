package movementtypes

import (
	"encoding/json"
	"net/http"
)

func HdlrGetMovementTypes(w http.ResponseWriter, r *http.Request) {
	movementTypes, err := ServGetMovementTypes()
	if err != nil {
		http.Error(w, "Error retrieving movement types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movementTypes)
}
