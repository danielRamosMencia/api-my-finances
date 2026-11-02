package movementtypes

func ServGetMovementTypes() ([]MovementTypeDTO, error) {
	movementTypes, err := RepoGetMovementTypes()
	if err != nil {
		return nil, err
	}

	var dtos []MovementTypeDTO

	for _, mt := range movementTypes {
		dtos = append(dtos, toDTO(mt))
	}

	return dtos, nil
}

func toDTO(movementType MovementTypeDB) MovementTypeDTO {
	return MovementTypeDTO{
		Pid:         movementType.Pid,
		Active:      movementType.Active,
		Name:        movementType.Name,
		Description: movementType.Description,
		Code:        movementType.Code,
	}
}
