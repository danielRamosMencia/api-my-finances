package configurations

import (
	"log/slog"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type EnvVariable struct {
	SERVER_PORT    string
	DATABASE_URL   string
	DATABASE_TOKEN string
	HMAC_KEY       string
}

var (
	once sync.Once
	Envs EnvVariable
)

func InitEnvVariables() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}

		Envs.SERVER_PORT = os.Getenv("SERVER_PORT")
		Envs.DATABASE_URL = os.Getenv("DATABASE_URL")
		Envs.DATABASE_TOKEN = os.Getenv("DATABASE_TOKEN")
		Envs.HMAC_KEY = os.Getenv("HMAC_KEY")

		slog.Info("Environment variables initialized")
	})
}
