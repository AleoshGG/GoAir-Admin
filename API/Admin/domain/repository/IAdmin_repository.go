package repository

import "GoAir-Admin/API/Admin/domain/entities"

type IAdmin interface {
	GetAdmin() entities.Admin
	CreatePlace(name string, id_user int) (uint, error)
	SearchUser(last_name string) entities.User 
	CreateId(id_place int) (error)
	GetIds(id_place int) []entities.Sensor
	GetPlaces(id_user int) []entities.Place
	DeletePlace(id_place int) (uint, error)
}