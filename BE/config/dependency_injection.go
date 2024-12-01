package config

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/controller"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/repository"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/service"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/utility"
	"gorm.io/gorm"
)

func UserDependencyInjection(db *gorm.DB) controller.UserController {
	userRepository := repository.NewUserRepository(db)
	jwtUtils := utility.NewJWTUtils()
	userService := service.NewUserService(userRepository, jwtUtils)
	userController := controller.NewUserController(userService)

	return userController
}
