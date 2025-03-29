package controllers

import (
	"GoAir-Admin/API/Admin/application/services"
	usecases "GoAir-Admin/API/Admin/application/useCases"
	"GoAir-Admin/API/Admin/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeletePlaceController struct {
	app  *usecases.DeletePlace
	auth *services.Auth
}

func NewDeletePlaceController() *DeletePlaceController {
	postgres := infrastructure.GetPostgreSQL()
	jwt := infrastructure.GetJWT()
	app := usecases.NewDeletePlace(postgres)
	auth := services.NewAuth(jwt)
	return &DeletePlaceController{app: app, auth: auth}
}

func (dp_c *DeletePlaceController) DeletePalce(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	id := c.Param("id")
	id_place, _ := strconv.ParseInt(id, 10, 64)

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	_, err := dp_c.auth.Run(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	rowsAffected, _ := dp_c.app.Run(int(id_place))
	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "No se pudo eliminar: No se entontró la referencia o ocurrió algo más",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Recurso eliminado",
	})
}