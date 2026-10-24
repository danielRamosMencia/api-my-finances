package users

func ServGetAll() ([]UserDTO, error) {
	users, err := RepoGetAll()
	if err != nil {
		return nil, err
	}

	var dtos []UserDTO

	for i := 0; i < len(users); i++ {
		dtos = append(dtos, toDTO(users[i]))
	}

	return dtos, nil
}

func toDTO(user UserDB) UserDTO {
	return UserDTO{
		Pid:       user.Pid,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Type:      user.Type,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
