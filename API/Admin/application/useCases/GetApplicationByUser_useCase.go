package usecases

import (
	"GoAir-Admin/API/Admin/domain/entities"
	"GoAir-Admin/API/Admin/domain/repository"
)

type GetApplicationByUser struct {
	db repository.IAdmin
}

func NewGetApplicationByUser(db repository.IAdmin) *GetApplicationByUser {
	return &GetApplicationByUser{db: db}
}

func (uc *GetApplicationByUser) Run(id_user int) []entities.Application {
	return uc.db.GetApplicationByUser(id_user)
}