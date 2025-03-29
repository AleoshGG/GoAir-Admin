package controllers

import (
	"GoAir-Admin/API/Admin/application/services"
	usecases "GoAir-Admin/API/Admin/application/useCases"
	"GoAir-Admin/API/Admin/infrastructure"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchUserController struct {
	app *usecases.SearchUser
	auth *services.Auth
}

func NewSearhUserController() *SearchUserController {
	postgres := infrastructure.GetPostgreSQL()
	jwt := infrastructure.GetJWT()
	app := usecases.NewSearchUser(postgres)
	auth := services.NewAuth(jwt)
	return &SearchUserController{app: app, auth: auth}
}

func (su_c *SearchUserController) SearchUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	last_name := c.Query("last_name")

	fmt.Println(tokenString)
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	fmt.Println(tokenString)
	_, err := su_c.auth.Run(tokenString)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	user := su_c.app.Run(last_name)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/users/",
		},
		"User": user,
	})


}

