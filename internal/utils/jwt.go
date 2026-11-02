package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/danielRamosMencia/api-my-finances/internal/configurations"
	"github.com/golang-jwt/jwt/v5"
)

const (
	JWT_ISSUER   = "api-my-finances"
	JWT_DURATION = time.Hour * 24
)

type claimsKey struct{}

type TokenRequest struct {
	Name     string
	Username string
	Pid      string
	Email    string
}

type TokenResponse struct {
	Token     string
	Session   string
	ExpiresAt time.Time
}

func ContextWithClaims(ctx context.Context, claims jwt.MapClaims) context.Context {
	return context.WithValue(ctx, claimsKey{}, claims)
}

func GetClaim(ctx context.Context, claim string) string {
	claims, ok := ctx.Value(claimsKey{}).(jwt.MapClaims)
	if !ok {
		return ""
	}

	value, _ := claims[claim].(string)
	return value
}

func jwtKey() []byte {
	return []byte(configurations.Envs.JWT_SECRET)
}

func CreateJwt(request TokenRequest) (TokenResponse, error) {
	now := time.Now()
	expiresAt := now.Add(JWT_DURATION)
	session := RandomString(32)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":      JWT_ISSUER,
			"name":     request.Name,
			"username": request.Username,
			"pid":      request.Pid,
			"email":    request.Email,
			"session":  session,
			"iat":      now.Unix(),
			"exp":      expiresAt.Unix(),
		})

	tokenString, err := token.SignedString(jwtKey())
	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{
		Token:     tokenString,
		Session:   session,
		ExpiresAt: expiresAt,
	}, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) {
			return jwtKey(), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer(JWT_ISSUER),
		jwt.WithExpirationRequired(),
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
