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

func HdlrUpdate(w http.ResponseWriter, r *http.Request) {
	var pid *string

	pid = utils.Ptr(r.PathValue("pid"))
	if pid == nil {
		http.Error(w, "not provided pid", http.StatusBadRequest)
		return
	}

	var request UserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ServUpdate(*pid, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

func HdlrUpdateActive(w http.ResponseWriter, r *http.Request) {
	var pid *string

	pid = utils.Ptr(r.PathValue("pid"))
	if pid == nil {
		http.Error(w, "not provided pid", http.StatusBadRequest)
		return
	}

	var activeRequest UpdateActiveRequest

	err := json.NewDecoder(r.Body).Decode(&activeRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ServUpdateActive(*pid, activeRequest.Active)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

func HdlrUpdatePassword(w http.ResponseWriter, r *http.Request) {
	var pid *string

	pid = utils.Ptr(r.PathValue("pid"))
	if pid == nil {
		http.Error(w, "not provided pid", http.StatusBadRequest)
		return
	}

	var passwordRequest UpdatePasswordRequest

	err := json.NewDecoder(r.Body).Decode(&passwordRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ServUpdatePassword(*pid, passwordRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func HdlrUpdateToken(w http.ResponseWriter, r *http.Request) {
	var pid *string

	pid = utils.Ptr(r.PathValue("pid"))
	if pid == nil {
		http.Error(w, "not provided pid", http.StatusBadRequest)
		return
	}

	token, err := ServUpdateToken(*pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(token)
}

func HdlrVerifyToken(w http.ResponseWriter, r *http.Request) {
	var pid *string

	pid = utils.Ptr(r.PathValue("pid"))
	if pid == nil {
		http.Error(w, "not provided pid", http.StatusBadRequest)
		return
	}

	var request VerifyTokenRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	valid, err := ServVerifyToken(*pid, request.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(valid)
}
