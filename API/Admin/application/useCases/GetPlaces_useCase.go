package usecases

import (
	"GoAir-Admin/API/Admin/domain/entities"
	"GoAir-Admin/API/Admin/domain/repository"
)

type GetPlaces struct {
	db repository.IAdmin
}

func NewGetPlaces(db repository.IAdmin) *GetPlaces {
	return &GetPlaces{db: db}
}

func (uc *GetPlaces) Run(id_user int) []entities.Place {
	return uc.db.GetPlaces(id_user)
}