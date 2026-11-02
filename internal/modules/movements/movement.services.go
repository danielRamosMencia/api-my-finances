package movements

import "github.com/danielRamosMencia/api-my-finances/internal/utils"

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

func toDTO(movement MovementDB) MovementDTO {
	return MovementDTO{
		Pid:              movement.Pid,
		Active:           movement.Active,
		Amount:           movement.Amount,
		AccountPid:       movement.AccountPid,
		AccountAlias:     movement.AccountAlias,
		MovementTypePid:  movement.MovementTypePid,
		MovementTypeName: movement.MovementTypeName,
		UserPid:          movement.UserPid,
		Username:         movement.Username,
		CreatedAt:        movement.CreatedAt,
	}
}
