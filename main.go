package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/configurations"
)

func main() {
	configurations.InitEnvVariables()
	serverPort := configurations.Envs.SERVER_PORT

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health/check", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello world")
	})

	slog.Info("Server listening on http://127.0.1.1" + serverPort)

	http.ListenAndServe(serverPort, mux)
}
