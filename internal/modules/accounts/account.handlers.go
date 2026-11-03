package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
)

func HldrGetUserAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := ServGetUserAccounts(r.Context())
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(accounts)
}

func HldrGetByPid(w http.ResponseWriter, r *http.Request) {
	pid := r.PathValue("pid")
	if pid == "" {
		responses.Error(w, http.StatusBadRequest, responses.MISSING_PID)
		return
	}

	account, err := ServGetByPid(pid)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(account)
}

func HldrSave(w http.ResponseWriter, r *http.Request) {
	var request AccountRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.InvalidBody(w, err)
		return
	}

	account, err := ServSave(r.Context(), request)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(account)
}

func HldrUpdate(w http.ResponseWriter, r *http.Request) {
	pid := r.PathValue("pid")
	if pid == "" {
		responses.Error(w, http.StatusBadRequest, responses.MISSING_PID)
		return
	}

	var request AccountRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.InvalidBody(w, err)
		return
	}

	account, err := ServUpdate(pid, request)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(account)
}
