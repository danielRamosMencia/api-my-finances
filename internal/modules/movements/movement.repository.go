package movements

import (
	"context"
	"log/slog"

	"github.com/danielRamosMencia/api-my-finances/internal/tursodb"
)

func RepoGetUserMovements(userPid string) ([]MovementDB, error) {
	var movements []MovementDB

	rows, err := tursodb.DB.QueryContext(
		context.Background(),
		SELECT_USER_MOVEMENTS,
		userPid,
	)

	if err != nil {
		slog.Error("[RepoGetUserMovements] error getting user movements: " + err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var movement MovementDB
		err := rows.Scan(
			&movement.Id,
			&movement.Pid,
			&movement.Active,
			&movement.Amount,
			&movement.AccountId,
			&movement.AccountPid,
			&movement.AccountAlias,
			&movement.MovementTypeId,
			&movement.MovementTypePid,
			&movement.MovementTypeName,
			&movement.UserId,
			&movement.UserPid,
			&movement.Username,
			&movement.CreatedAt,
			&movement.UpdatedAt,
			&movement.DeletedAt,
		)
		if err != nil {
			slog.Error("[RepoGetUserMovements] error scanning movement: " + err.Error())
			return nil, err
		}
		movements = append(movements, movement)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[RepoGetUserMovements] error checking movements: " + err.Error())
		return nil, err
	}

	return movements, nil
}
