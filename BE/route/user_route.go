package route

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.Engine, userController controller.UserController) {
	userRoute := route.Group("/api/user")
	{
		userRoute.POST("/register", userController.Register)
	}
}
