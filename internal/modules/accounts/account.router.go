package accounts

import "net/http"

func SetAccountsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/accounts/users", HldrGetUserAccounts)
	mux.HandleFunc("GET /api/accounts/{pid}", HldrGetByPid)
	mux.HandleFunc("POST /api/accounts", HldrSave)
	mux.HandleFunc("PUT /api/accounts/{pid}", HldrUpdate)
}
