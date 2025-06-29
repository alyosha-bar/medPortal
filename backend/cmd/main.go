// @title           medPortal API
// @version         1.0
// @description     This is my Receptionist's and Doctor's API in Golang

// @contact.name   Aleksej Barysnikov
// @contact.url    https://aleksejbarysnikov.netlify.app/
// @contact.email  alohahoy@gmail.com

// @host      localhost:8080
// @BasePath  /

package main

import (
	"time"

	"github.com/alyosha-bar/medPortal/database"
	"github.com/alyosha-bar/medPortal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/alyosha-bar/medPortal/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// connect to DB
	database.ConnectDB()

	// Set Up GIN router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://diligent-perception-production.up.railway.app", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Serve Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(router)

	// start server
	router.Run(":8080")
}
