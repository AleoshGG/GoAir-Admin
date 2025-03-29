package main

import (
	iAdmin "GoAir-Admin/API/Admin/infrastructure"
	Aroutes "GoAir-Admin/API/Admin/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	iAdmin.GoDependences()
	
	r := gin.Default()
	r.Use(cors.Default())

	Aroutes.RegisterRouter(r)
	
	r.Run(":8020")
}