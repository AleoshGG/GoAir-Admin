package infrastructure

import "GoAir-Admin/API/Admin/infrastructure/adapters"

var postgres *adapters.PostgreSQL
var JWT *adapters.JWT
var rabbitmq *adapters.RabbitMQ

func GoDependences() {
	postgres = adapters.NewPostgreSQL()
	JWT = adapters.NewJWT()
	rabbitmq = adapters.NewRabbitMQ()
}

func GetPostgreSQL() *adapters.PostgreSQL {
	return postgres
}

func GetJWT() *adapters.JWT {
	return JWT
}

func GetRabbitMQ() *adapters.RabbitMQ {
	return rabbitmq
}