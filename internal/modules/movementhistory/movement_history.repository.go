package movementhistory

import (
	"context"
	"log/slog"

	"github.com/danielRamosMencia/api-my-finances/internal/tursodb"
)

func RepoGetUserHistory(userPid string) ([]MovementHistoryDB, error) {
	history := []MovementHistoryDB{}

	rows, err := tursodb.DB.QueryContext(
		context.Background(),
		SELECT_USER_HISTORY,
		userPid,
	)

	if err != nil {
		slog.Error("[RepoGetUserHistory] error getting user movement history: " + err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var item MovementHistoryDB
		err := rows.Scan(
			&item.Pid,
			&item.HistoryAmount,
			&item.Amount,
			&item.MovementCode,
			&item.Account,
			&item.Type,
			&item.TypeCode,
			&item.Username,
			&item.Email,
			&item.Date,
		)

		if err != nil {
			slog.Error("[RepoGetUserHistory] errror scannig history: " + err.Error())
			return nil, err
		}

		history = append(history, item)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[RepoGetUserHistory] error checking history: " + err.Error())
		return nil, err
	}

	return history, err
}
