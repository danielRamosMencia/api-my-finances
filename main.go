package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/danielRamosMencia/api-my-finances/internal/configurations"
	"github.com/danielRamosMencia/api-my-finances/internal/middlewares"
	"github.com/danielRamosMencia/api-my-finances/internal/modules/accounts"
	"github.com/danielRamosMencia/api-my-finances/internal/modules/auth"
	"github.com/danielRamosMencia/api-my-finances/internal/modules/health"
	"github.com/danielRamosMencia/api-my-finances/internal/modules/movements"
	"github.com/danielRamosMencia/api-my-finances/internal/modules/movementtypes"
	"github.com/danielRamosMencia/api-my-finances/internal/modules/users"
	"github.com/danielRamosMencia/api-my-finances/internal/sessions"
	"github.com/danielRamosMencia/api-my-finances/internal/tursodb"
)

func main() {
	configurations.InitEnvVariables()
	tursodb.ConnectTursoDb()
	sessions.StartCleanup(15 * time.Minute)

	mux := http.NewServeMux()

	auth.SetAuthRoutes(mux)
	health.SetHealthRoutes(mux)
	users.SetUserRoutes(mux)
	accounts.SetAccountsRoutes(mux)
	movementtypes.SetMovementTypeRoutes(mux)
	movements.SetMovementRoutes(mux)

	handler := middlewares.AuthRequired(mux,
		"POST /api/auth/login",
		"PATCH /api/users/{pid}/password",
		"PATCH /api/users/{pid}/token",
		"GET /api/health",
		"GET /api/mock-pid",
	)

	slog.Info("Server listening on http://127.0.1.1" + configurations.Envs.SERVER_PORT)

	http.ListenAndServe(configurations.Envs.SERVER_PORT, handler)
}
