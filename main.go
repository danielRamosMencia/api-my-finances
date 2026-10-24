package main

import (
	"log/slog"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/configurations"
	"github.com/danielRamosMencia/api-my-finances/internal/modules/users"
	"github.com/danielRamosMencia/api-my-finances/internal/tursodb"
)

func main() {
	configurations.InitEnvVariables()
	tursodb.ConnectTursoDb()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/users", users.HdlrGetAll)

	slog.Info("Server listening on http://127.0.1.1" + configurations.Envs.SERVER_PORT)

	http.ListenAndServe(configurations.Envs.SERVER_PORT, mux)
}
