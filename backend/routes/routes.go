package routes

import (
	"github.com/alyosha-bar/medPortal/handlers"
	"github.com/alyosha-bar/medPortal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	authRoutes := router.Group("/auth")
	{
		// login
		authRoutes.POST("/login", handlers.Login)

		// sign up
		authRoutes.POST("/signup", handlers.SignUp)

	}

	// following use auth middleware functions
	// receptionists' routes
	receptionistRoutes := router.Group("/receptionist")
	receptionistRoutes.Use(middleware.AuthMiddleware("receptionist"))
	{

		// GET ALL patients
		receptionistRoutes.GET("/patients", handlers.GetAllPatients)

		// GET SPECIFIC patient
		receptionistRoutes.GET("/patient/:patient_id", handlers.GetPatient)

		// CREATE NEW patient
		receptionistRoutes.POST("/register", handlers.RegisterPatient)

		// Update patient details --> simple changes
		receptionistRoutes.PATCH("/details/update")

		// UPDATE patient --> assign to doctor --> need to change DB schema
		receptionistRoutes.PATCH("/patient/assign")

		// Delete patient profile
		receptionistRoutes.DELETE("/patient/:patient_id", handlers.DeletePatientProfile)
	}

	// doctors' routes
	doctorsRoutes := router.Group("/doctor")
	doctorsRoutes.Use(middleware.AuthMiddleware("doctor"))
	{
		// GET patients which belong to the doctor
		doctorsRoutes.GET("/patients")

		// GET patient details
		doctorsRoutes.GET("/patient/:patient_id")

		// UPDATE patient record --> medical notes

	}

}
