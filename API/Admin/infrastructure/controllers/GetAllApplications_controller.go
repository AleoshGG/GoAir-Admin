package controllers

import (
	"GoAir-Admin/API/Admin/application/services"
	usecases "GoAir-Admin/API/Admin/application/useCases"
	"GoAir-Admin/API/Admin/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllApplicationsController struct {
	app *usecases.GetAllApplications
	auth *services.Auth
}

func NewGetAllApplications() *GetAllApplicationsController {
	postgres := infrastructure.GetPostgreSQL()
	jwt      := infrastructure.GetJWT()
	app := usecases.NewGetAllApplications(postgres)
	auth := services.NewAuth(jwt)
	return &GetAllApplicationsController{app: app, auth: auth}
} 

func (gaa_c *GetAllApplicationsController) GetAllApplications(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	_, err := gaa_c.auth.Run(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	apps := gaa_c.app.Run()
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/admin/",
		},
		"data": apps,
	})
}

