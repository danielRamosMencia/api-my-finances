package auth

import (
	"context"
	"log/slog"

	"github.com/danielRamosMencia/api-my-finances/internal/tursodb"
)

func RepoGetUserByUsername(username string) (*UserDB, error) {
	row := tursodb.DB.QueryRowContext(
		context.Background(),
		SELECT_USER_BY_USERNAME,
		username,
	)

	var user UserDB

	err := row.Scan(
		&user.Name,
		&user.Username,
		&user.Pid,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		slog.Error("[RepoGetUserByUsername] error getting user by username: " + err.Error())
		return nil, err
	}

	return &user, nil
}
