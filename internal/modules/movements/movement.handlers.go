package movements

import (
	"encoding/json"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/middlewares"
)

func HldrGetUserMovements(w http.ResponseWriter, r *http.Request) {
	ok := middlewares.AuthRequired(w, r)
	if !ok {
		return
	}

	movements, err := ServGetUserMovements()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movements)
}

func HldrGetByPid(w http.ResponseWriter, r *http.Request) {
	pid := r.PathValue("pid")
	if pid == "" {
		http.Error(w, "not provided pid", http.StatusBadRequest)
		return
	}

	movement, err := ServGetByPid(pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := ServSave(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(account)
}
