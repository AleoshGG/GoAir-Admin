package usecases

import (
	"GoAir-Admin/API/Admin/domain/entities"
	"GoAir-Admin/API/Admin/domain/repository"
)

type SearchUser struct {
	db repository.IAdmin
}

func NewSearchUser(db repository.IAdmin) *SearchUser {
	return &SearchUser{db: db}
}

func (uc SearchUser) Run(last_name string) entities.User {
	return uc.db.SearchUser(last_name)
}