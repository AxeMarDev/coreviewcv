package main

import (
	"backendcv/internal/database"
	"backendcv/internal/routes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {

	database.Db = database.InitDB()

	fmt.Println("started")
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS", "PATCH", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "credentials"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.Routes(router)

	log.Println("Starting HTTPS server on https://localhost:443...")
	err := router.RunTLS(":443", "/etc/letsencrypt/live/railiant-appservice.app/fullchain.pem", "/etc/letsencrypt/live/railiant-appservice.app/privkey.pem")
	if err != nil {

		log.Println("Starting HTTP server on https://localhost:8080...")
		err = router.Run("localhost:8080")

		if err != nil {
			fmt.Println("fatal error")
		}
	}
}
