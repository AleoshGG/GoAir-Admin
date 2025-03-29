package usecases

import (
	"GoAir-Admin/API/Admin/domain/entities"
	"GoAir-Admin/API/Admin/domain/repository"
)
type GetIds struct {
	db repository.IAdmin
}

func NewGetIds(db repository.IAdmin) *GetIds {
	return &GetIds{db: db}
}

func (uc *GetIds) Run(id_place int) []entities.Sensor {
	return uc.db.GetIds(id_place)
}