package utils

import (
	"fmt"
	"time"

	"github.com/danielRamosMencia/api-my-finances/internal/configurations"
	"github.com/golang-jwt/jwt/v5"
)

func MockGetClaim(claim string) string {
	// TODO: Implement a mock function to retrieve claims from a JWT token for testing purposes.
	return "nsPTs1LNq1zk5zRV1OkimiDSocGYSyfX"
}

var JWT_KEY = []byte(configurations.Envs.JWT_SECRET)

type TokenRequest struct {
	Name     string
	Username string
	Pid      string
	Email    string
}

func CreateJwt(request TokenRequest) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":      "api-my-finances",
			"name":     request.Name,
			"username": request.Username,
			"pid":      request.Pid,
			"email":    request.Email,
			"session":  RandomString(16),
			"iat":      time.Now().Unix(),
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(JWT_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
