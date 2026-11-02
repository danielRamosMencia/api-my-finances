package movementtypes

import (
	"context"
	"log/slog"

	"github.com/danielRamosMencia/api-my-finances/internal/tursodb"
)

func RepoGetMovementTypes() ([]MovementTypeDB, error) {
	rows, err := tursodb.DB.QueryContext(
		context.Background(),
		SELECT_MOVEMENT_TYPES,
	)
	if err != nil {
		slog.Error("[RepoGetMovementTypes] error querying movement types: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	var movementTypes []MovementTypeDB

	for rows.Next() {
		var movementType MovementTypeDB
		err := rows.Scan(
			&movementType.Pid,
			&movementType.Active,
			&movementType.Name,
			&movementType.Code,
			&movementType.Description,
		)

		if err != nil {
			slog.Error("[RepoGetMovementTypes] error scanning movement type: " + err.Error())
			return nil, err
		}

		movementTypes = append(movementTypes, movementType)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[RepoGetMovementTypes] error checking movement types: " + err.Error())
		return nil, err
	}

	return movementTypes, nil
}
