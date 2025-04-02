package controllers

import (
	"GoAir-Admin/API/Admin/application/services"
	usecases "GoAir-Admin/API/Admin/application/useCases"
	"GoAir-Admin/API/Admin/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetIdsSensorsController struct {
	app *usecases.GetIds
	auth *services.Auth
}

func NewGetIdsSensorsController() *GetIdsSensorsController {
	postgres := infrastructure.GetPostgreSQL()
	jwt := infrastructure.GetJWT()
	app := usecases.NewGetIds(postgres)
	auth := services.NewAuth(jwt)
	return &GetIdsSensorsController{app: app, auth: auth}
}

func (gids_c *GetIdsSensorsController) GetIds(c *gin.Context) {
	id := c.Param("id")
	id_place, _ := strconv.ParseInt(id, 10, 64)

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	_, err := gids_c.auth.Run(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	sensors, devices := gids_c.app.Run(int(id_place))

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/users/",
		},
		"Sensors": sensors,
		"Devices": devices,
		"id_place": id_place,
	})


}