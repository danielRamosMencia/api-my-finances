package health

import "net/http"

func SetHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/health", HdlrHealthCheck)
}
