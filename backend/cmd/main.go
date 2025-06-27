package main

import (
	"github.com/alyosha-bar/medPortal/database"
	"github.com/alyosha-bar/medPortal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to DB
	database.ConnectDB()

	// initialise middleware

	// Set Up GIN router
	router := gin.Default()
	routes.SetupRoutes(router)

	// start server
	router.Run(":8080")
}
