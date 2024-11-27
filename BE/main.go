package main

import (
	"fmt"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/command"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/config"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/constants"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/middleware"
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

	server := gin.Default()
	server.Use(middleware.CORS())

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	var env = os.Getenv("APP_ENV")
	fmt.Printf("Environment is %s\n", env)
	var serve string
	if env == constants.ENUM_ENV_PRODUCTION {
		serve = ":" + port
	} else if env == constants.ENUM_ENV_DEVELOPMENT {
		serve = "localhost:" + port
	}
	err = server.Run(serve)

	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
