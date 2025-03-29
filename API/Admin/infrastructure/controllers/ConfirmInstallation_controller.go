package controllers

import (
	"GoAir-Admin/API/Admin/application/services"
	usecases "GoAir-Admin/API/Admin/application/useCases"
	"GoAir-Admin/API/Admin/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConfirmInstallationController struct {
	app  *usecases.ConfirmInstallation
	auth *services.Auth
}

func NewConfirmInstallationController() *ConfirmInstallationController {
	postgres := infrastructure.GetPostgreSQL()
	jwt := infrastructure.GetJWT()
	app := usecases.NewConfirmInstallation(postgres)
	auth := services.NewAuth(jwt)
	return &ConfirmInstallationController{app: app, auth: auth}
}

func (ci_c *ConfirmInstallationController) ConfirmInstallation(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	id := c.Param("id")
	id_application, _ := strconv.ParseInt(id, 10, 64)

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	_, err := ci_c.auth.Run(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	err = ci_c.app.Run(int(id_application))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se puedo actualizar", "error:":err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Estado actualizado",
	})
}