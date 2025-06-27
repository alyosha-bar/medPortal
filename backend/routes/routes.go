package routes

import (
	"github.com/alyosha-bar/medPortal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	authRoutes := router.Group("/auth")
	{
		// login ???
		authRoutes.POST("/login", handlers.Login)

		// sign up ???
		authRoutes.POST("/signup", handlers.SignUp)

	}

	// following use auth middleware functions
	// receptionists' routes
	receptionistRoutes := router.Group("/receptionist")
	{

		// GET ALL patients
		receptionistRoutes.GET("/patients")

		// GET SPECIFIC patient

		// CREATE NEW patient

		// Update patient details

		// Delete patient profile
	}

	// doctors' routes

}
