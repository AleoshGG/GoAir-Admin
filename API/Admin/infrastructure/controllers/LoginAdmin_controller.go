package controllers

import (
	"GoAir-Admin/API/Admin/application/services"
	usecases "GoAir-Admin/API/Admin/application/useCases"
	"GoAir-Admin/API/Admin/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginAdminController struct {
	app         *usecases.GetAdmin
	jwt 		*services.CreateJWT
}

func NewLoginAdminController() *LoginAdminController {
	postgres := infrastructure.GetPostgreSQL()
	jwt := infrastructure.GetJWT()
	app := usecases.NewGetAdmin(postgres)
	jwts := services.NewCreateJWT(jwt)
	return &LoginAdminController{app: app, jwt: jwts}
}

func (l_c *LoginAdminController) Login(c *gin.Context) {
	var credentials struct {
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Datos inválidos: " + err.Error(),
		})
		return
	}

	admin := l_c.app.Run()
	if admin.Email != credentials.Email || admin.Password != credentials.Password {
		c.JSON(http.StatusForbidden, gin.H{
			"status": false,
			"error":  "Credenciales incorrectas: ",
		})
		return
	}

	token, err := l_c.jwt.Run(admin)
	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  "Error al generar el JWT: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/admin/",
		},
		"token": token,
	})

}