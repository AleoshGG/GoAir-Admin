package services

import (
	"GoAir-Admin/API/Admin/domain/entities"
)

type Auth struct {
	jwt IServices
}

func NewAuth(jwt IServices) *Auth {
	return &Auth{jwt: jwt}
}

func (t *Auth) Run(token string) (entities.Claims, error) {
	return t.jwt.Auth(token)
}