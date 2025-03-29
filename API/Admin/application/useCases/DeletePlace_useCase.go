package usecases

import (
	"GoAir-Admin/API/Admin/domain/repository"
)

type DeletePlace struct {
	db repository.IAdmin
}

func NewDeletePlace(db repository.IAdmin) *DeletePlace {
	return &DeletePlace{db: db}
}

func (uc DeletePlace) Run(id_place int) (uint, error) {
	return uc.db.DeletePlace(id_place)
}