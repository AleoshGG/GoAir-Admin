package controllers

import (
	"GoAir-Admin/API/Admin/application/services"
	usecases "GoAir-Admin/API/Admin/application/useCases"
	"GoAir-Admin/API/Admin/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetApplicationByUserController struct {
	app *usecases.GetApplicationByUser
	auth *services.Auth
}

func NewGetApplicationByUser() *GetApplicationByUserController {
	postgres := infrastructure.GetPostgreSQL()
	jwt      := infrastructure.GetJWT()
	app := usecases.NewGetApplicationByUser(postgres)
	auth := services.NewAuth(jwt)
	return &GetApplicationByUserController{app: app, auth: auth}
} 

func (gabu_c *GetApplicationByUserController) GetApplicationByUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	id := c.Param("id")
	id_user, _ := strconv.ParseInt(id, 10, 64)

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	_, err := gabu_c.auth.Run(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	apps := gabu_c.app.Run(int(id_user))
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/admin/",
		},
		"data": apps,
	})
}

