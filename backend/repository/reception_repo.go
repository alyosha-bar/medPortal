package repository

import (
	"github.com/alyosha-bar/medPortal/database"
	"github.com/alyosha-bar/medPortal/models"
)

func GetAllPatients() ([]models.Patient, error) {
	var patients []models.Patient
	result := database.DB.Find(&patients)
	return patients, result.Error
}

func GetPatient(patientID uint) (models.Patient, error) {
	var patient models.Patient
	result := database.DB.Where("id = ?", patientID).First(&patient)
	return patient, result.Error
}

func RegisterPatient(patient models.Patient) (models.Patient, error) {
	result := database.DB.Create(&patient)
	return patient, result.Error
}

func DeletePatientProfile(patientID uint) error {
	result := database.DB.Delete(&models.User{}, patientID)
	return result.Error
}
