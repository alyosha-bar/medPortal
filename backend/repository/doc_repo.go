package repository

import (
	"github.com/alyosha-bar/medPortal/database"
	"github.com/alyosha-bar/medPortal/models"
)

func GetPatientsByDoctor(doctorID uint) ([]models.Patient, error) {
	var patients []models.Patient
	result := database.DB.Where("doctor_id = ?", doctorID).Find(&patients)
	return patients, result.Error
}
