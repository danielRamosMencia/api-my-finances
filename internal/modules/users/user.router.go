package users

import (
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/modules/accounts"
	"github.com/danielRamosMencia/api-my-finances/internal/modules/movementhistory"
	"github.com/danielRamosMencia/api-my-finances/internal/modules/movements"
)

func SetUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/users", HdlrGetAll)
	mux.HandleFunc("GET /api/users/{pid}", HdlrGetByPid)
	mux.HandleFunc("POST /api/users", HdlrSave)
	mux.HandleFunc("PUT /api/users/{pid}", HdlrUpdate)
	mux.HandleFunc("PATCH /api/users/{pid}/active", HdlrUpdateActive)
	mux.HandleFunc("PATCH /api/users/{pid}/password", HdlrUpdatePassword)
	mux.HandleFunc("PATCH /api/users/{pid}/token", HdlrUpdateToken)
	mux.HandleFunc("PATCH /api/users/{pid}/token/verify", HdlrVerifyToken)

	// accounts
	mux.HandleFunc("GET /api/users/accounts", accounts.HldrGetUserAccounts)

	// movements
	mux.HandleFunc("GET /api/users/movements", movements.HldrGetUserMovements)
	mux.HandleFunc("GET /api/users/movements/history", movementhistory.HldrGetUserHistory)
}
