package tursodb

import (
	"context"
	"database/sql"
	"log/slog"
	"sync"

	"github.com/danielRamosMencia/api-my-finances/internal/configurations"
	turso "turso.tech/database/tursogo-serverless"
)

var (
	once sync.Once
	DB   *sql.DB
)

func ConnectTursoDb() {
	once.Do(func() {
		DB := sql.OpenDB(turso.NewConnector(
			configurations.Envs.DATABASE_URL,
			configurations.Envs.DATABASE_TOKEN,
		))
		defer DB.Close()

		ping, err := DB.QueryContext(context.Background(), "SELECT '1' as ping;")
		if err != nil {
			panic(err)
		}
		if ping.Err() != nil {
			panic(ping.Err())
		}

		slog.Info("Remote Turso database connected")
	})
}
