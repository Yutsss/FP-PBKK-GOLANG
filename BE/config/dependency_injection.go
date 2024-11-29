package config

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/controller"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/repository"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/service"
	"gorm.io/gorm"
)

func UserDependencyInjection(db *gorm.DB) controller.UserController {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	return userController
}
