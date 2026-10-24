package users

import (
	"context"
	"log"

	"github.com/danielRamosMencia/api-my-finances/internal/tursodb"
)

func RepoGetAll() ([]UserDB, error) {
	rows, err := tursodb.DB.QueryContext(context.Background(), SELECT_ALL)
	if err != nil {
		log.Printf("[RepoGetAll] error selecting users: %v", err)
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
			return nil, err
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
