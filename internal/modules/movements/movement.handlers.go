package movements

import (
	"encoding/json"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
)

func HldrGetUserMovements(w http.ResponseWriter, r *http.Request) {
	movements, err := ServGetUserMovements(r.Context())
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(movements)
}

func HldrGetByPid(w http.ResponseWriter, r *http.Request) {
	pid := r.PathValue("pid")
	if pid == "" {
		responses.Error(w, http.StatusBadRequest, responses.MISSING_PID)
		return
	}

	movement, err := ServGetByPid(pid)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(movement)
}

func HldrSave(w http.ResponseWriter, r *http.Request) {
	var request MovementRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.InvalidBody(w, err)
		return
	}

	movement, err := ServSave(r.Context(), request)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(movement)
}
