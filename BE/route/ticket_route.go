package route

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/controller"
	"github.com/gin-gonic/gin"
)

func TicketRouter(route *gin.Engine, ticketController controller.TicketController, middlewares map[string]gin.HandlerFunc) {
	ticketRoute := route.Group("/api/ticket")
	{
		ticketRoute.POST("/create", middlewares["authMiddleware"], middlewares["userRoleMiddleware"], ticketController.Create)
		ticketRoute.GET("/all", middlewares["authMiddleware"], middlewares["adminRoleMiddleware"], ticketController.GetAll)
		ticketRoute.GET("/user", middlewares["authMiddleware"], middlewares["userRoleMiddleware"], ticketController.GetByUserId)
		ticketRoute.GET("/:ticket_id", middlewares["authMiddleware"], middlewares["allRoleMiddleware"], ticketController.GetById)
		ticketRoute.PUT("/:ticket_id/assign/:technician_id", middlewares["authMiddleware"], middlewares["adminRoleMiddleware"], ticketController.AssignById)
	}
}
