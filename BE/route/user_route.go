package route

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.Engine, userController controller.UserController, middlewares map[string]gin.HandlerFunc) {
	userRoute := route.Group("/api/user")
	{
		userRoute.POST("/register", userController.Register)
		userRoute.POST("/login", userController.Login)
		userRoute.GET("/me", middlewares["authMiddleware"], middlewares["allRoleMiddleware"], userController.Get)
		userRoute.GET("/all", middlewares["authMiddleware"], middlewares["adminRoleMiddleware"], userController.GetAll)
		userRoute.POST("/logout", middlewares["authMiddleware"], middlewares["allRoleMiddleware"], userController.Logout)
	}
}
