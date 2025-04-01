package routes

import (
	"GoAir-Admin/API/Admin/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	adminRoutes := r.Group("/admin")
	{
		adminRoutes.POST("/login", controllers.NewLoginAdminController().Login)
		adminRoutes.GET("/search", controllers.NewSearhUserController().SearchUser)
		adminRoutes.GET("/places/:id", controllers.NewGetPlacesUserController().GetPlacesUser)
		adminRoutes.GET("/sensors/:id", controllers.NewGetIdsSensorsController().GetIds)
		adminRoutes.POST("/", controllers.NewCreatePlaceController().CreatePlace)
		adminRoutes.DELETE("/:id", controllers.NewDeletePlaceController().DeletePalce)
		adminRoutes.GET("/apps/:id", controllers.NewGetApplicationByUser().GetApplicationByUser)
		adminRoutes.GET("/apps", controllers.NewGetAllApplications().GetAllApplications)
		adminRoutes.DELETE("/confirm/:id", controllers.NewConfirmInstallationController().ConfirmInstallation)
	}
}

