package movements

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/modules/movementtypes"
	"github.com/danielRamosMencia/api-my-finances/internal/responses"
	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

func ServGetUserMovements(ctx context.Context) ([]MovementDTO, error) {
	userPid := utils.GetClaim(ctx, "pid")
	movements, err := RepoGetUserMovements(userPid)

	if err != nil {
		return nil, err
	}

	var dtos []MovementDTO
	for _, movement := range movements {
		dtos = append(dtos, toDTO(movement))
	}

	return dtos, nil
}

func ServGetByPid(pid string) (MovementDTO, error) {
	movement, err := RepoGetByPid(pid)
	if err != nil {
		return MovementDTO{}, err
	}

	if movement == nil {
		return MovementDTO{}, responses.NewError(http.StatusNotFound, responses.MOVEMENT_NOT_FOUND)
	}

	return toDTO(*movement), nil
}

func ServSave(ctx context.Context, request MovementRequest) (MovementDTO, error) {
	typeCode, err := RepoGetMovementTypeCode(request.MovementTypePid)
	if errors.Is(err, sql.ErrNoRows) {
		return MovementDTO{}, responses.NewError(http.StatusNotFound, responses.MOVEMENT_TYPE_NOT_FOUND)
	}

	if err != nil {
		return MovementDTO{}, err
	}

	income := typeCode == movementtypes.INCOME_CODE

	instant := utils.GetNowFormat()
	userPid := utils.GetClaim(ctx, "pid")
	code := utils.RandomNumber(6, true)
	pid := utils.RandomString(32)

	movement := MovementDB{
		Pid:             pid,
		Active:          true,
		Code:            code,
		Description:     request.Description,
		Amount:          request.Amount,
		AccountPid:      request.AccountPid,
		MovementTypePid: request.MovementTypePid,
		UserPid:         userPid,
		CreatedAt:       instant,
		UpdatedAt:       utils.Ptr(instant),
	}

	err = RepoSave(movement, income)
	if err != nil {
		return MovementDTO{}, err
	}

	return toDTO(movement), nil
}

func toDTO(movement MovementDB) MovementDTO {
	return MovementDTO{
		Pid:              movement.Pid,
		Active:           movement.Active,
		Amount:           movement.Amount,
		Code:             movement.Code,
		Description:      movement.Description,
		AccountPid:       movement.AccountPid,
		AccountAlias:     movement.AccountAlias,
		MovementTypePid:  movement.MovementTypePid,
		MovementTypeName: movement.MovementTypeName,
		UserPid:          movement.UserPid,
		Username:         movement.Username,
		CreatedAt:        movement.CreatedAt,
	}
}
