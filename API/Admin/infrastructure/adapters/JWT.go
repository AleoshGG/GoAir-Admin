package adapters

import (
	"GoAir-Admin/API/Admin/domain/entities"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type JWT struct{}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) CreateJWT(admin entities.Admin) (string, error) {
	claims := entities.Claims{
		Password: admin.Password,
		Email:   admin.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Expira en 2 horas
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   fmt.Sprintf("%s", admin.Email),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (j *JWT) Auth(tokenString string) (entities.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &entities.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return entities.Claims{}, err
	}

	claims, ok := token.Claims.(*entities.Claims)

	if !ok || !token.Valid {
		return entities.Claims{}, fmt.Errorf("token inválido")
	}

	return *claims, nil
}