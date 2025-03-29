package usecases

import (
	"GoAir-Admin/API/Admin/domain/repository"
)

type CreatePlace struct {
	db repository.IAdmin
}

func NewCreatePlace(db repository.IAdmin) *CreatePlace {
	return &CreatePlace{db: db}
}

func (uc CreatePlace) Run(name string, id_user int, id_application int) (uint, error) {
	return uc.db.CreatePlace(name, id_user, id_application)
}