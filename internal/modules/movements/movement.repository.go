package movements

import (
	"context"
	"log/slog"

	"github.com/danielRamosMencia/api-my-finances/internal/tursodb"
	"github.com/danielRamosMencia/api-my-finances/internal/utils"
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
			&movement.Code,
			&movement.Description,
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

func RepoGetByPid(pid string) (*MovementDB, error) {
	row := tursodb.DB.QueryRowContext(
		context.Background(),
		SELECT_BY_PID,
		pid,
	)

	var movement MovementDB

	err := row.Scan(
		&movement.Id,
		&movement.Pid,
		&movement.Code,
		&movement.Description,
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
		slog.Error("[RepoGetByPid] error getting movement by pid: " + err.Error())
		return nil, err
	}

	return &movement, nil
}

func RepoSave(movement MovementDB, income bool) error {
	ctx := context.Background()

	// Begin transaction
	tx, err := tursodb.DB.BeginTx(context.Background(), nil)
	if err != nil {
		slog.Error("[RepoSave] error starting transaction: " + err.Error())
		return err
	}

	defer tx.Rollback()

	var accountId int
	account := tx.QueryRowContext(
		ctx,
		SELECT_ACCOUNT_ID,
		movement.AccountPid,
		movement.UserPid,
	)

	err = account.Scan(&accountId)
	if err != nil {
		slog.Error("[RepoSave] error getting account for movement: " + err.Error())
		return err
	}

	var movementTypeId int
	movementType := tx.QueryRowContext(
		ctx,
		SELECT_MOVEMENT_TYPE_ID,
		movement.MovementTypePid,
	)

	err = movementType.Scan(&movementTypeId)
	if err != nil {
		slog.Error("[RepoSave] error getting type for movement: " + err.Error())
		return err
	}

	var userId int
	user := tx.QueryRowContext(
		ctx,
		SELECT_USER_ID,
		movement.UserPid,
	)

	err = user.Scan(&userId)
	if err != nil {
		slog.Error("[RepoSave] error getting type for movement: " + err.Error())
		return err
	}

	result, err := tx.ExecContext(
		ctx,
		INSERT,
		movement.Pid,
		movement.Active,
		movement.Code,
		movement.Description,
		movement.Amount,
		accountId,
		movementTypeId,
		userId,
		movement.CreatedAt,
		movement.UpdatedAt,
	)

	if err != nil {
		slog.Error("[RepoSave] error saving movement: " + err.Error())
		return err
	}

	var amount float64
	if income {
		amount = float64(movement.Amount * -1)
	} else {
		amount = movement.Amount
	}

	_, err = tx.ExecContext(
		ctx,
		UPDATE_ACCOUNT_AMOUNT,
		amount,
		movement.UpdatedAt,
		accountId,
	)

	if err != nil {
		slog.Error("[RepoSave] error updating account amount: " + err.Error())
		return err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		slog.Error("[RepoSave] error extracting last inserted ID: " + err.Error())
	}

	_, err = tx.ExecContext(
		ctx,
		INSERT_MOVEMENT_HISTORY,
		utils.RandomString(32),
		movement.Amount,
		lastId,
		movement.CreatedAt,
	)

	if err != nil {
		slog.Error("[RepoSave]: error saving movement history: " + err.Error())
		return err
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		slog.Error("[RepoSave] error committing transaction: " + err.Error())
		return err
	}

	tursodb.ShowDbStatistics()

	return nil
}
