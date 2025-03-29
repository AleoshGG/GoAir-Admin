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

func (uc CreatePlace) Run(name string, id_user int) (uint, error) {
	return uc.db.CreatePlace(name, id_user)
}