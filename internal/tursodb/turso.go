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

		DB.SetMaxOpenConns(3)
		DB.SetMaxIdleConns(0)
		DB.SetConnMaxLifetime(0)

		slog.Info("Remote Turso database connected")
	})
}

func CloseTursoDb() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

func ShowDbStatistics() {
	stats := DB.Stats()

	slog.Debug("Database connection pool:",
		"open_connections", stats.OpenConnections,
		"in_use", stats.InUse,
		"idle", stats.Idle,
		"wait_count", stats.WaitCount,
		"wait_duration", stats.WaitDuration,
	)
}
