package configurations

import (
	"log/slog"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type EnvVariable struct {
	SERVER_PORT         string
	DATABASE_URL        string
	DATABASE_TOKEN      string
	HMAC_KEY            string
	JWT_SECRET          string
	TRUST_PROXY_HEADERS bool
}

var (
	once sync.Once
	Envs EnvVariable
)

func InitEnvVariables() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			slog.Warn("No .env file found, using system environment variables")
		}

		Envs.SERVER_PORT = os.Getenv("SERVER_PORT")
		Envs.DATABASE_URL = os.Getenv("DATABASE_URL")
		Envs.DATABASE_TOKEN = os.Getenv("DATABASE_TOKEN")
		Envs.HMAC_KEY = os.Getenv("HMAC_KEY")
		Envs.JWT_SECRET = os.Getenv("JWT_SECRET")
		Envs.TRUST_PROXY_HEADERS = os.Getenv("TRUST_PROXY_HEADERS") == "true"

		slog.Info("Environment variables initialized")
	})
}
