package health

import "net/http"

func SetHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/health", HdlrHealthCheck)
	mux.HandleFunc("GET /api/mock-pid", HldrMockPid)
	mux.HandleFunc("GET /api/health/sessions", HldrDevSessions)
	mux.HandleFunc("GET /api/health/rate-limits", HldrDevRateLimits)
}
