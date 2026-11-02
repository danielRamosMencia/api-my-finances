package health

import (
	"encoding/json"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/middlewares"
	"github.com/danielRamosMencia/api-my-finances/internal/sessions"
	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

func HdlrHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"status":"ok"}`))
}

func HldrMockPid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	pid := utils.RandomString(32)
	json.NewEncoder(w).Encode(pid)
}

// Development only: remove before production
func HldrDevSessions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(sessions.Snapshot())
}

// Development only: remove before production
func HldrDevRateLimits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(middlewares.LimitsSnapshot())
}
