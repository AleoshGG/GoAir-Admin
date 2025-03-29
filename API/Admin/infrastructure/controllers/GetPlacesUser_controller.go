package controllers

import (
	"GoAir-Admin/API/Admin/application/services"
	usecases "GoAir-Admin/API/Admin/application/useCases"
	"GoAir-Admin/API/Admin/infrastructure"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetPlacesUserController struct {
	app *usecases.GetPlaces
	auth *services.Auth
}

func NewGetPlacesUserController() *GetPlacesUserController {
	postgres := infrastructure.GetPostgreSQL()
	jwt := infrastructure.GetJWT()
	app := usecases.NewGetPlaces(postgres)
	auth := services.NewAuth(jwt)
	return &GetPlacesUserController{app: app, auth: auth}
}

func (gpu_c *GetPlacesUserController) GetPlacesUser(c *gin.Context) {
	id := c.Param("id")
	id_user, _ := strconv.ParseInt(id, 10, 64)

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	_, err := gpu_c.auth.Run(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	places := gpu_c.app.Run(int(id_user))
	fmt.Print(places)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/users/",
		},
		"places": places,
	})


}