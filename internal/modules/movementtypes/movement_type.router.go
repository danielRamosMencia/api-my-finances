package movementtypes

import "net/http"

func SetMovementTypeRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/movement-types", HdlrGetMovementTypes)
}
