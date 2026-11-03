package users

import (
	"encoding/json"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

func HdlrGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := ServGetAll()

	if err != nil {
		responses.HandleError(w, err)
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
		responses.Error(w, http.StatusBadRequest, responses.MISSING_PID)
		return
	}

	user, err := ServGetByPid(*pid)
	if err != nil {
		responses.HandleError(w, err)
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
		responses.InvalidBody(w, err)
		return
	}

	user, err := ServSave(request)
	if err != nil {
		responses.HandleError(w, err)
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
		responses.Error(w, http.StatusBadRequest, responses.MISSING_PID)
		return
	}

	var request UserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.InvalidBody(w, err)
		return
	}

	user, err := ServUpdate(*pid, request)
	if err != nil {
		responses.HandleError(w, err)
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
		responses.Error(w, http.StatusBadRequest, responses.MISSING_PID)
		return
	}

	var activeRequest UpdateActiveRequest

	err := json.NewDecoder(r.Body).Decode(&activeRequest)
	if err != nil {
		responses.InvalidBody(w, err)
		return
	}

	user, err := ServUpdateActive(*pid, activeRequest.Active)
	if err != nil {
		responses.HandleError(w, err)
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
		responses.Error(w, http.StatusBadRequest, responses.MISSING_PID)
		return
	}

	var passwordRequest UpdatePasswordRequest

	err := json.NewDecoder(r.Body).Decode(&passwordRequest)
	if err != nil {
		responses.InvalidBody(w, err)
		return
	}

	err = ServUpdatePassword(*pid, passwordRequest.Password)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func HdlrUpdateToken(w http.ResponseWriter, r *http.Request) {
	var pid *string

	pid = utils.Ptr(r.PathValue("pid"))
	if pid == nil {
		responses.Error(w, http.StatusBadRequest, responses.MISSING_PID)
		return
	}

	token, err := ServUpdateToken(*pid)
	if err != nil {
		responses.HandleError(w, err)
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
		responses.Error(w, http.StatusBadRequest, responses.MISSING_PID)
		return
	}

	var request VerifyTokenRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.InvalidBody(w, err)
		return
	}

	valid, err := ServVerifyToken(*pid, request.Token)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(valid)
}
