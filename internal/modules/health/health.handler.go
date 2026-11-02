package health

import (
	"encoding/json"
	"net/http"

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
