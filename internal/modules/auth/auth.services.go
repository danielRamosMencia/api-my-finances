package auth

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
	"github.com/danielRamosMencia/api-my-finances/internal/sessions"
	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

func ServLogin(request LoginRequest) (LoginResponse, error) {
	invalidCredentials := responses.NewError(http.StatusUnauthorized, responses.INVALID_CREDENTIALS)

	user, err := RepoGetUserByUsername(request.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return LoginResponse{}, invalidCredentials
	}

	if err != nil {
		return LoginResponse{}, err
	}

	match := utils.CheckPasswordHash(request.Password, user.Password)
	if !match {
		return LoginResponse{}, invalidCredentials
	}

	jwt, err := utils.CreateJwt(utils.TokenRequest{
		Name:     user.Name,
		Username: user.Username,
		Pid:      user.Pid,
		Email:    user.Email,
	})
	if err != nil {
		return LoginResponse{}, err
	}

	sessions.Save(jwt.Session, jwt.Token, jwt.ExpiresAt)

	return LoginResponse{
		Username: user.Username,
		Pid:      user.Pid,
		Name:     user.Name,
		Email:    user.Email,
		Token:    jwt.Token,
	}, nil
}

func ServLogout(ctx context.Context) {
	sessions.Delete(utils.GetClaim(ctx, "session"))
}
