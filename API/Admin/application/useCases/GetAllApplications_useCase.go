package usecases

import (
	"GoAir-Admin/API/Admin/domain/entities"
	"GoAir-Admin/API/Admin/domain/repository"
)

type GetAllApplications struct {
	db repository.IAdmin
}

func NewGetAllApplications(db repository.IAdmin) *GetAllApplications {
	return &GetAllApplications{db: db}
}

func (uc *GetAllApplications) Run() []entities.AllApplications {
	return uc.db.GetAllApplications()
}