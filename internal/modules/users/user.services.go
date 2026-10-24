package users

import (
	"time"

	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

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

func ServSave(request UserRequest) (UserDTO, error) {
	db := UserDB{
		Id:        1,
		Pid:       utils.RandomString(32),
		Active:    true,
		Name:      request.Name,
		Username:  request.Username,
		Email:     request.Email,
		Token:     utils.Ptr(utils.RandomString(8)),
		Password:  request.Password,
		Type:      "OPERATOR",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: utils.Ptr(time.Now().UTC()),
		DeletedAt: utils.Ptr(time.Now().UTC()),
	}

	return toDTO(db), nil
}

func toDTO(user UserDB) UserDTO {
	return UserDTO{
		Pid:       user.Pid,
		Name:      user.Name,
		Active:    user.Active,
		Username:  user.Username,
		Email:     user.Email,
		Type:      user.Type,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
