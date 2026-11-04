package movementhistory

import (
	"context"
	"log/slog"

	"github.com/danielRamosMencia/api-my-finances/internal/tursodb"
)

func RepoGetUserHistory(userPid, from, to string) ([]MovementHistoryDB, error) {
	history := []MovementHistoryDB{}

	rows, err := tursodb.DB.QueryContext(
		context.Background(),
		SELECT_USER_HISTORY,
		userPid,
		from,
		to,
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

func RepoGetUserMetrics(userPid, from, to string) (UserMetrics, error) {
	ctx := context.Background()
	metrics := UserMetrics{
		Accounts: []AccountMetrics{},
	}

	row := tursodb.DB.QueryRowContext(
		ctx,
		SELECT_USER_METRICS,
		userPid,
		from,
		to,
	)

	err := row.Scan(
		&metrics.TotalMovements,
		&metrics.TotalIncome,
		&metrics.TotalExpense,
		&metrics.Net,
	)

	if err != nil {
		slog.Error("[RepoGetUserMetrics] error getting user metrics: " + err.Error())
		return UserMetrics{}, err
	}

	rows, err := tursodb.DB.QueryContext(
		ctx,
		SELECT_USER_ACCOUNT_METRICS,
		userPid,
		from,
		to,
	)

	if err != nil {
		slog.Error("[RepoGetUserMetrics] error getting user account metrics: " + err.Error())
		return UserMetrics{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var account AccountMetrics
		err := rows.Scan(
			&account.AccountPid,
			&account.Account,
			&account.TotalMovements,
			&account.TotalIncome,
			&account.TotalExpense,
			&account.Net,
		)

		if err != nil {
			slog.Error("[RepoGetUserMetrics] error scanning account metrics: " + err.Error())
			return UserMetrics{}, err
		}

		metrics.Accounts = append(metrics.Accounts, account)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[RepoGetUserMetrics] error checking account metrics: " + err.Error())
		return UserMetrics{}, err
	}

	return metrics, nil
}
