package controllers

import (
	"GoAir-Admin/API/Admin/application/services"
	usecases "GoAir-Admin/API/Admin/application/useCases"
	"GoAir-Admin/API/Admin/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePlaceController struct {
	app *usecases.CreatePlace
	auth *services.Auth
}

func NewCreatePlaceController() *CreatePlaceController {
	postgres := infrastructure.GetPostgreSQL()
	jwt := infrastructure.GetJWT()
	app := usecases.NewCreatePlace(postgres)
	auth := services.NewAuth(jwt)
	return &CreatePlaceController{app: app, auth: auth}
}

func (cp_c *CreatePlaceController) CreatePlace(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	var newPlace struct {
		Id_user    int
		Name string
	}

	if err := c.ShouldBindJSON(&newPlace); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Datos inválidos: " + err.Error(),
		})
		return
	}

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	_, err := cp_c.auth.Run(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	id_place, err := cp_c.app.Run(newPlace.Name, newPlace.Id_user) 
	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  "Error al crear el nuevo espacio " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/admin/",
		},
		"id_place": id_place,
	})


}