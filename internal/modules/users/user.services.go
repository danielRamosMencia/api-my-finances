package users

import (
	"errors"
	"time"

	"github.com/danielRamosMencia/api-my-finances/internal/configurations"
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

func ServUpdate(pid string, request UserRequest) (UserDTO, error) {
	user, err := RepoGetByPid(pid)
	if err != nil {
		return UserDTO{}, err
	}

	if user == nil {
		return UserDTO{}, errors.New("user with pid " + pid + " not found")
	}

	instant := utils.GetNowFormat()

	user.Name = request.Name
	user.Username = request.Username
	user.Email = request.Email
	user.UpdatedAt = utils.Ptr(instant)

	err = RepoUpdate(pid, *user)
	if err != nil {
		return UserDTO{}, err
	}

	return toDTO(*user), nil
}

func ServUpdateActive(pid string, active bool) (UserDTO, error) {
	user, err := RepoGetByPid(pid)
	if err != nil {
		return UserDTO{}, err
	}

	if user == nil {
		return UserDTO{}, errors.New("user with pid " + pid + " not found")
	}

	err = RepoUpdateActive(pid, active)
	if err != nil {
		return UserDTO{}, err
	}

	user.Active = active

	return toDTO(*user), nil
}

func ServUpdatePassword(pid string, password string) error {
	user, err := RepoGetByPid(pid)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user with pid " + pid + " not found")
	}

	hashedPassword, _ := utils.HashPassword(password)

	err = RepoUpdatePassword(pid, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func ServUpdateToken(pid string) (string, error) {
	user, err := RepoGetByPid(pid)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user with pid " + pid + " not found")
	}

	newToken := utils.RandomString(8)
	token := utils.HmacSHA256(configurations.Envs.HMAC_KEY, newToken)

	err = RepoUpdateToken(pid, token)
	if err != nil {
		return "", err
	}

	return newToken, nil
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
