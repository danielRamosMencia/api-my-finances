package auth

import (
	"encoding/json"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
)

func HldrLogin(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.InvalidBody(w, err)
		return
	}

	response, err := ServLogin(request)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func HldrLogout(w http.ResponseWriter, r *http.Request) {
	ServLogout(r.Context())

	w.WriteHeader(http.StatusNoContent)
}
