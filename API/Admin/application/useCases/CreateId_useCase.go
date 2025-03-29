package usecases

import (
	"GoAir-Admin/API/Admin/domain/repository"
)

type CreateId struct {
	db repository.IAdmin
}

func NewCreateId(db repository.IAdmin) *CreateId {
	return &CreateId{db: db}
}

func (uc *CreateId) Run(id_place int) (error) {
	return uc.db.CreateId(id_place)
}