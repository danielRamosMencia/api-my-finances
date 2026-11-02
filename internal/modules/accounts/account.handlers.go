package accounts

import (
	"encoding/json"
	"net/http"
)

func HldrGetUserAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := ServGetUserAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(accounts)
}

func HldrGetByPid(w http.ResponseWriter, r *http.Request) {
	pid := r.PathValue("pid")
	if pid == "" {
		http.Error(w, "not provided pid", http.StatusBadRequest)
		return
	}

	account, err := ServGetByPid(pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

func HldrUpdate(w http.ResponseWriter, r *http.Request) {
	pid := r.PathValue("pid")
	if pid == "" {
		http.Error(w, "not provided pid", http.StatusBadRequest)
		return
	}

	var request AccountRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := ServUpdate(pid, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(account)
}
