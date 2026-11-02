package users

import "net/http"

func SetUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/users", HdlrGetAll)
	mux.HandleFunc("GET /api/users/{pid}", HdlrGetByPid)
	mux.HandleFunc("POST /api/users", HdlrSave)
}
