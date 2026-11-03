package movementhistory

import (
	"encoding/json"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
)

func HldrGetUserHistory(w http.ResponseWriter, r *http.Request) {
	history, err := ServGetUserHistory(r.Context())
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(history)
}
