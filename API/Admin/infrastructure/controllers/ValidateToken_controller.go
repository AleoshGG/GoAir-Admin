package controllers

import (
	"GoAir-Admin/API/Admin/application/services"
	"GoAir-Admin/API/Admin/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidateTokenController struct {

	auth *services.Auth
}

func NewValidateTokenController() *ValidateTokenController {
	jwt := infrastructure.GetJWT()
	auth := services.NewAuth(jwt)
	return &ValidateTokenController{auth: auth}
}

func (dp_c *ValidateTokenController) ValidateToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

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


	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Token Válido",
	})
}