package main

import (
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
	err := godotenv.Load(".env")

	db := config.ConnectDB()
	defer config.CloseDBConnection(db)

	if len(os.Args) > 1 {
		flag := command.Commands(db)
		if !flag {
			return
		}
	}

	if err != nil {
		log.Fatalf("Error loading env file")
	}

	var env = os.Getenv("APP_ENV")
	var serve string

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	if env == constants.ENUM_ENV_DEVELOPMENT {
		gin.SetMode(gin.DebugMode)
		serve = "localhost:" + port
	} else if env == constants.ENUM_ENV_PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
		serve = ":" + port
	} else {
		panic("Invalid APP_ENV")
	}

	server := gin.Default()
	server.Use(middleware.CORS())

	middlewares := config.MiddlewareDependencyInjection()
	userController := config.UserDependencyInjection(db)

	route.UserRouter(server, userController, middlewares)

	err = server.Run(serve)

	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
