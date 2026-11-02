package movements

import (
	"errors"

	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

func ServGetUserMovements() ([]MovementDTO, error) {
	userPid := utils.MockGetClaim("pid")
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
		return MovementDTO{}, errors.New("movement with pid " + pid + " not found")
	}

	return toDTO(*movement), nil
}

func ServSave(request MovementRequest) (MovementDTO, error) {
	instant := utils.GetNowFormat()
	userPid := utils.MockGetClaim("pid")
	code := utils.RandomNumber(6, true)
	pid := utils.RandomString(32)

	movement := MovementDB{
		Pid:             pid,
		Active:          true,
		Code:            code,
		Amount:          request.Amount,
		AccountPid:      request.AccountPid,
		MovementTypePid: request.MovementTypePid,
		UserPid:         userPid,
		CreatedAt:       instant,
		UpdatedAt:       utils.Ptr(instant),
	}

	err := RepoSave(movement, request.Income)
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
		AccountPid:       movement.AccountPid,
		AccountAlias:     movement.AccountAlias,
		MovementTypePid:  movement.MovementTypePid,
		MovementTypeName: movement.MovementTypeName,
		UserPid:          movement.UserPid,
		Username:         movement.Username,
		CreatedAt:        movement.CreatedAt,
	}
}
