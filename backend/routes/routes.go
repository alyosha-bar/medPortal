package routes

import (
	"github.com/alyosha-bar/medPortal/database"
	"github.com/alyosha-bar/medPortal/handlers"
	"github.com/alyosha-bar/medPortal/middleware"
	"github.com/alyosha-bar/medPortal/repository"
	"github.com/alyosha-bar/medPortal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// Initialize repositories
	receptionRepo := repository.NewReceptionRepo(database.DB)
	doctorRepo := repository.NewDoctorRepo(database.DB)
	authRepo := repository.NewAuthRepo(database.DB)

	// Initialize services
	receptionService := services.NewReceptionService(receptionRepo)
	doctorService := services.NewDoctorService(doctorRepo)
	authService := services.NewAuthService(authRepo)

	// Initialize handlers
	receptionHandler := handlers.NewReceptionHandler(receptionService)
	doctorHandler := handlers.NewDoctorHandler(doctorService)
	authHandler := handlers.NewAuthHandler(authService)

	authRoutes := router.Group("/auth")
	{
		// login
		authRoutes.POST("/login", authHandler.Login)

		// sign up
		authRoutes.POST("/signup", authHandler.SignUp)

	}

	// following use auth middleware functions
	// receptionists' routes
	receptionistRoutes := router.Group("/receptionist")
	receptionistRoutes.Use(middleware.AuthMiddleware("receptionist"))
	{

		// GET ALL patients
		receptionistRoutes.GET("/patients", receptionHandler.GetAllPatients)

		// GET SPECIFIC patient
		receptionistRoutes.GET("/patient/:patient_id", receptionHandler.GetPatient)

		// CREATE NEW patient
		receptionistRoutes.POST("/register", receptionHandler.RegisterPatient)

		// Update patient details --> simple changes to profile
		receptionistRoutes.PATCH("/details/update/:patient_id", receptionHandler.UpdateField)

		// UPDATE patient --> assign to doctor
		receptionistRoutes.PATCH("/patient/assign/:patient_id", receptionHandler.AssignPatient)

		// GET all doctor names --> used for the endpoint above feature
		receptionistRoutes.GET("/doctors", receptionHandler.GetAllDoctors)

		// Delete patient profile
		receptionistRoutes.DELETE("/patient/:patient_id", receptionHandler.DeletePatientProfile)
	}

	// doctors' routes
	doctorsRoutes := router.Group("/doctor")
	doctorsRoutes.Use(middleware.AuthMiddleware("doctor"))
	{
		// GET patients which belong to the doctor
		doctorsRoutes.GET("/patients", doctorHandler.GetPatientsByDoctor)

		// UPDATE patient record --> medical notes
		doctorsRoutes.PATCH("/patient/:patient_id", doctorHandler.UpdateMedicalNotes)

	}
}
