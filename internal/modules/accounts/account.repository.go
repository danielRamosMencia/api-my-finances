package accounts

import (
	"context"
	"log/slog"

	"github.com/danielRamosMencia/api-my-finances/internal/tursodb"
)

func RepoGetUserAccounts(userPid string) ([]AccountDB, error) {
	accounts := []AccountDB{}

	rows, err := tursodb.DB.QueryContext(
		context.Background(),
		SELECT_BY_USER_PID,
		userPid,
	)

	if err != nil {
		slog.Error("[RepoGetUserAccounts] error getting user accounts: " + err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var account AccountDB
		err := rows.Scan(
			&account.Id,
			&account.Pid,
			&account.Active,
			&account.Alias,
			&account.amount,
			&account.userId,
			&account.userPid,
			&account.CreatedAt,
			&account.UpdatedAt,
		)

		if err != nil {
			slog.Error("[RepoGetUserAccounts] error scanning account: " + err.Error())
			return nil, err
		}

		accounts = append(accounts, account)
	}
	if err := rows.Err(); err != nil {
		slog.Error("[RepoGetUserAccounts] error checking accounts: " + err.Error())
		return nil, err
	}

	return accounts, nil
}

func RepoGetByPid(pid string) (*AccountDB, error) {
	row := tursodb.DB.QueryRowContext(
		context.Background(),
		SELECT_BY_PID,
		pid,
	)

	var account AccountDB

	err := row.Scan(
		&account.Id,
		&account.Pid,
		&account.Active,
		&account.Alias,
		&account.amount,
		&account.userId,
		&account.userPid,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		slog.Error("[RepoGetByPid] error getting account by pid: " + err.Error())
		return nil, err
	}

	return &account, nil
}

func RepoInsert(account AccountDB) error {
	_, err := tursodb.DB.ExecContext(
		context.Background(),
		INSERT,
		account.Pid,
		account.Alias,
		account.amount,
		account.userPid,
	)

	if err != nil {
		slog.Error("[RepoInsert] error inserting account: " + err.Error())
		return err
	}

	return nil
}

func RepoUpdate(pid string, account AccountDB) error {
	_, err := tursodb.DB.ExecContext(
		context.Background(),
		UPDATE,
		account.Alias,
		account.amount,
		account.UpdatedAt,
		pid,
	)

	if err != nil {
		slog.Error("[RepoUpdate] error updating account: " + err.Error())
		return err
	}

	return nil
}
