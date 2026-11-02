package users

import (
	"errors"
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

func ServGetByPid(pid string) (UserDTO, error) {
	user, err := RepoGetByPid(pid)
	if err != nil {
		return UserDTO{}, err
	}

	if user == nil {
		return UserDTO{}, errors.New("user with pid " + pid + " not found")
	}

	return toDTO(*user), nil
}

func ServSave(request UserRequest) (UserDTO, error) {
	hashedPassword, _ := utils.HashPassword(request.Password)
	instant := time.Now().UTC().Format(time.RFC3339Nano)

	user := UserDB{
		Pid:       utils.RandomString(32),
		Active:    true,
		Name:      request.Name,
		Username:  request.Username,
		Email:     request.Email,
		Token:     utils.Ptr(utils.RandomString(8)),
		Password:  hashedPassword,
		Type:      request.Type,
		CreatedAt: instant,
		UpdatedAt: utils.Ptr(instant),
	}

	err := RepoSave(user)
	if err != nil {
		return UserDTO{}, err
	}

	return toDTO(user), nil
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
