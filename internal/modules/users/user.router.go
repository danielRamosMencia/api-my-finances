package users

import "net/http"

func SetUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/users", HdlrGetAll)
	mux.HandleFunc("GET /api/users/{pid}", HdlrGetByPid)
	mux.HandleFunc("POST /api/users", HdlrSave)
	mux.HandleFunc("PUT /api/users/{pid}", HdlrUpdate)
	mux.HandleFunc("PATCH /api/users/{pid}/active", HdlrUpdateActive)
	mux.HandleFunc("PATCH /api/users/{pid}/password", HdlrUpdatePassword)
	mux.HandleFunc("PATCH /api/users/{pid}/token", HdlrUpdateToken)
	mux.HandleFunc("PATCH /api/users/{pid}/token/verify", HdlrVerifyToken)
}
