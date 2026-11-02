package users

import (
	"encoding/json"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/utils"
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

func HdlrSave(w http.ResponseWriter, r *http.Request) {
	var request UserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ServSave(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

func HdlrGetByPid(w http.ResponseWriter, r *http.Request) {
	var pid *string

	pid = utils.Ptr(r.PathValue("pid"))
	if pid == nil {
		http.Error(w, "not provided pid", http.StatusBadRequest)
		return
	}

	user, err := ServGetByPid(*pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
