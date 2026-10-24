package tursodb

import (
	"context"
	"database/sql"
	"log/slog"
	"sync"
	"time"

	"github.com/danielRamosMencia/api-my-finances/internal/configurations"
	turso "turso.tech/database/tursogo-serverless"
)

var (
	once sync.Once
	DB   *sql.DB
)

func ConnectTursoDb() {
	once.Do(func() {
		DB = sql.OpenDB(turso.NewConnector(
			configurations.Envs.DATABASE_URL,
			configurations.Envs.DATABASE_TOKEN,
		))

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := DB.PingContext(ctx); err != nil {
			panic(err)
		}

		slog.Info("Remote Turso database connected")
	})
}

func CloseTursoDb() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
