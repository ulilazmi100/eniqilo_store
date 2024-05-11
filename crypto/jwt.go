package crypto

import (
	"time"

	configs "eniqilo_store/cfg"
	"eniqilo_store/db/entities"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id, phone, name string) (string, error) {
	secret := configs.GetEnvOrDefault("JWT_SECRET", "sec")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entities.JWTClaims{
		Id:    id,
		Phone: phone,
		Name:  name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	})

	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func VerifyToken(token string) (*entities.JWTPayload, error) {
	secret := configs.GetEnvOrDefault("JWT_SECRET", "sec")

	claims := &entities.JWTClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims.RegisteredClaims.ExpiresAt.Before(time.Now()) {
		return nil, err
	}

	payload := &entities.JWTPayload{
		Id:    claims.Id,
		Phone: claims.Phone,
		Name:  claims.Name,
	}

	return payload, nil
}
