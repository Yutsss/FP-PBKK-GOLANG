package main

import (
	"fmt"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/command"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/config"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/constants"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/middleware"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	var env string
	if os.Getenv("GO_ENV") == constants.ENUM_ENV_DEVELOPMENT {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading env file")
		}
		env = constants.ENUM_ENV_DEVELOPMENT
	} else if os.Getenv("GO_ENV") == constants.ENUM_ENV_PRODUCTION {
		err := godotenv.Load(".env.production")
		if err != nil {
			log.Fatalf("Error loading env file")
		}
		env = constants.ENUM_ENV_PRODUCTION
	} else {
		panic("Invalid GO_ENV")
	}

	fmt.Printf("Environment is %s\n", env)

	db := config.ConnectDB()
	defer config.CloseDBConnection(db)

	if len(os.Args) > 1 {
		flag := command.Commands(db)
		if !flag {
			return
		}
	}

	var serve string

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	switch env {
	case constants.ENUM_ENV_DEVELOPMENT:
		gin.SetMode(gin.DebugMode)
		serve = "localhost:" + port
	case constants.ENUM_ENV_PRODUCTION:
		gin.SetMode(gin.ReleaseMode)
		serve = ":" + port
	default:
		panic("Unexpected environment value")
	}

	server := gin.Default()
	server.Use(middleware.CORS())

	middlewares := config.MiddlewareDependencyInjection()
	userController := config.UserDependencyInjection(db)

	route.UserRouter(server, userController, middlewares)

	err := server.Run(serve)

	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
