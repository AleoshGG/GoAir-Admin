package usecases

import "GoAir-Admin/API/Admin/domain/repository"

type ConfirmInstallation struct {
	db repository.IAdmin
}

func NewConfirmInstallation(db repository.IAdmin) *ConfirmInstallation {
	return &ConfirmInstallation{db: db}
}

func (uc *ConfirmInstallation) Run(id_application int) (error) {
	return uc.db.ConfirmInstallation(id_application)
}