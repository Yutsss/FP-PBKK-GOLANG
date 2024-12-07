package config

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/constants"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/controller"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/middleware"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/repository"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/service"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/utility"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserDependencyInjection(db *gorm.DB) controller.UserController {
	userRepository := repository.NewUserRepository(db)
	jwtUtils := utility.NewJWTUtils()
	userService := service.NewUserService(userRepository, jwtUtils)
	userController := controller.NewUserController(userService)

	return userController
}

func TicketDependencyInjection(db *gorm.DB) controller.TicketController {
	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketController := controller.NewTicketController(ticketService)

	return ticketController
}

func MiddlewareDependencyInjection() map[string]gin.HandlerFunc {
	jwtUtils := utility.NewJWTUtils()

	middlewares := make(map[string]gin.HandlerFunc)

	middlewares["authMiddleware"] = middleware.AuthMiddleware(jwtUtils)
	middlewares["allRoleMiddleware"] = middleware.RoleMiddleware([]string{
		constants.ENUM_ROLE_ADMIN, constants.ENUM_ROLE_USER, constants.ENUM_ROLE_TECHNICIAN,
	})
	middlewares["userRoleMiddleware"] = middleware.RoleMiddleware([]string{constants.ENUM_ROLE_USER})
	middlewares["adminRoleMiddleware"] = middleware.RoleMiddleware([]string{constants.ENUM_ROLE_ADMIN})
	middlewares["technicianRoleMiddleware"] = middleware.RoleMiddleware([]string{constants.ENUM_ROLE_TECHNICIAN})

	return middlewares
}
