package auth

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/api-my-finances/internal/sessions"
	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

func ServLogin(request LoginRequest) (string, error) {
	user, err := RepoGetUserByUsername(request.Username)
	if err != nil {
		return "", err
	}

	match := utils.CheckPasswordHash(request.Password, user.Password)
	if !match {
		return "", errors.New("password not match")
	}

	jwt, err := utils.CreateJwt(utils.TokenRequest{
		Name:     user.Name,
		Username: user.Username,
		Pid:      user.Pid,
		Email:    user.Email,
	})
	if err != nil {
		return "", err
	}

	sessions.Save(jwt.Session, jwt.Token, jwt.ExpiresAt)

	return jwt.Token, nil
}

func ServLogout(ctx context.Context) {
	sessions.Delete(utils.GetClaim(ctx, "session"))
}
