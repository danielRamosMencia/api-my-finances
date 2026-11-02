package movements

import "net/http"

func SetMovementRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/movements/{pid}", HldrGetByPid)
	mux.HandleFunc("POST /api/movements", HldrSave)
}
