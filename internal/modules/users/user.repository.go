package users

import (
	"context"
	"log/slog"

	"github.com/danielRamosMencia/api-my-finances/internal/tursodb"
)

func RepoGetAll() ([]UserDB, error) {
	rows, err := tursodb.DB.QueryContext(context.Background(), SELECT_ALL)
	if err != nil {
		slog.Error("[RepoGetAll] error getting all users: " + err.Error())
		return nil, err
	}

	defer rows.Close()

	var users []UserDB

	for rows.Next() {
		var user UserDB
		if err := rows.Scan(
			&user.Pid,
			&user.Active,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Type,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			slog.Error("[RepoGetAll] error scannig all users: " + err.Error())
			return nil, err
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		slog.Error("[RepoGetAll] error checking all users: " + err.Error())
		return nil, err
	}

	return users, nil
}

func RepoSave(user UserDB) error {
	_, err := tursodb.DB.ExecContext(
		context.Background(),
		INSERT,
		user.Pid,
		user.Name,
		user.Username,
		user.Email,
		user.Token,
		user.Password,
		user.Type,
	)

	if err != nil {
		slog.Error("[RepoSave] error saving user: " + err.Error())
		return err
	}

	return nil
}
