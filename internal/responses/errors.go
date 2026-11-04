package responses

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

const (
	UNKNOWN_ERROR        = "UNKNOWN ERROR"
	ROUTE_NOT_FOUND      = "ROUTE NOT FOUND"
	RESOURCE_NOT_FOUND   = "RESOURCE NOT FOUND"
	INVALID_REQUEST_BODY = "INVALID REQUEST BODY"
	MISSING_PID          = "MISSING PID"
	INVALID_DATE_FILTER  = "INVALID DATE FILTER"

	INVALID_CREDENTIALS = "INVALID CREDENTIALS"
	MISSING_TOKEN       = "MISSING AUTHORIZATION TOKEN"
	INVALID_TOKEN       = "INVALID TOKEN"
	TOKEN_EXPIRED       = "TOKEN EXPIRED"
	SESSION_CLOSED      = "SESSION CLOSED"

	MISSING_DEVICE_ID                = "MISSING OR INVALID DEVICE ID"
	TOO_MANY_REQUESTS                = "TOO MANY REQUESTS"
	TOO_MANY_REQUESTS_SESSION_CLOSED = "TOO MANY REQUESTS, SESSION CLOSED"

	USER_NOT_FOUND     = "USER NOT FOUND"
	USER_HAS_NO_TOKEN  = "USER HAS NO TOKEN"
	ACCOUNT_NOT_FOUND  = "ACCOUNT NOT FOUND"
	MOVEMENT_NOT_FOUND = "MOVEMENT NOT FOUND"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
}

// AppError is an expected error whose message is safe to show to the client.
type AppError struct {
	Status  int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewError(status int, message string) *AppError {
	return &AppError{Status: status, Message: message}
}

func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorResponse{
		Message:   message,
		Timestamp: utils.GetNowFormat(),
		Status:    status,
	})
}

// HandleError responds with the AppError's status and message. Any other error is logged
// and the client only receives UNKNOWN ERROR.
func HandleError(w http.ResponseWriter, err error) {
	var appError *AppError
	if errors.As(err, &appError) {
		Error(w, appError.Status, appError.Message)
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		Error(w, http.StatusNotFound, RESOURCE_NOT_FOUND)
		return
	}

	slog.Error("[HandleError] unexpected error: " + err.Error())
	Error(w, http.StatusInternalServerError, UNKNOWN_ERROR)
}

func InvalidBody(w http.ResponseWriter, err error) {
	slog.Warn("[InvalidBody] error decoding request body: " + err.Error())
	Error(w, http.StatusBadRequest, INVALID_REQUEST_BODY)
}
