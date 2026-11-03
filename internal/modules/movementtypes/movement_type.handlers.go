package movementtypes

import (
	"encoding/json"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
)

func HdlrGetMovementTypes(w http.ResponseWriter, r *http.Request) {
	movementTypes, err := ServGetMovementTypes()
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movementTypes)
}
