package services

import "GoAir-Admin/API/Admin/domain/entities"

type CreateJWT struct {
	jwt IServices
}

func NewCreateJWT(jwt IServices) *CreateJWT {
	return &CreateJWT{jwt: jwt}
}

func (t *CreateJWT) Run(admin entities.Admin) (string, error) {
	return t.jwt.CreateJWT(admin)
}