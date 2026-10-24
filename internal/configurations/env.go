package configurations

import (
	"log/slog"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type EnvVariable struct {
	SERVER_PORT string
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

		slog.Info("Environment variables initialized")
	})
}
