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

		// Update patient details --> simple changes to profile
		receptionistRoutes.PATCH("/details/update/:patient_id", handlers.UpdateField)

		// UPDATE patient --> assign to doctor
		receptionistRoutes.PATCH("/patient/assign/:patient_id", handlers.AssignPatient)

		// GET all doctor names --> used for the endpoint above feature
		receptionistRoutes.GET("/doctors", handlers.GetAllDoctors)

		// Delete patient profile
		receptionistRoutes.DELETE("/patient/:patient_id", handlers.DeletePatientProfile)
	}

	// doctors' routes
	doctorsRoutes := router.Group("/doctor")
	doctorsRoutes.Use(middleware.AuthMiddleware("doctor"))
	{
		// GET patients which belong to the doctor
		doctorsRoutes.GET("/patients", handlers.GetPatientsByDoctor)

		// UPDATE patient record --> medical notes
		doctorsRoutes.PATCH("/patient/:patient_id", handlers.UpdateMedicalNotes)

	}
}
