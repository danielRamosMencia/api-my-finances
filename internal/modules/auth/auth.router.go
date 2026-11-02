package auth

import "net/http"

func SetAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/login", HldrLogin)
}
