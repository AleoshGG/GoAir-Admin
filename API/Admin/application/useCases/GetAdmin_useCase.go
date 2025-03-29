package usecases

import (
	"GoAir-Admin/API/Admin/domain/entities"
	"GoAir-Admin/API/Admin/domain/repository"
)

type GetAdmin struct {
	db repository.IAdmin
}

func NewGetAdmin(db repository.IAdmin) *GetAdmin {
	return &GetAdmin{db: db}
}

func (uc *GetAdmin) Run() entities.Admin {
	return uc.db.GetAdmin()
}