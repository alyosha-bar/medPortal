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
	result := database.DB.Delete(&models.Patient{}, patientID)
	return result.Error
}

func UpdateField(patientID uint, field string, value interface{}) (models.Patient, error) {
	var patient models.Patient

	// Perform update
	if err := database.DB.Model(&models.Patient{}).
		Where("id = ?", patientID).
		Update(field, value).Error; err != nil {
		return patient, err
	}

	// Return updated patient
	err := database.DB.First(&patient, patientID).Error
	return patient, err
}

func GetAllDoctors() ([]models.User, error) {
	var doctors []models.User
	result := database.DB.Select("id, username").Where("role = ?", "doctor").Find(&doctors)
	return doctors, result.Error
}

func AssignPatient(patientID uint, doctorID uint) (models.Patient, error) {
	var patient models.Patient

	result := database.DB.Model(&patient).
		Where("id = ?", patientID).
		Update("doctor_id", doctorID)

	if result.Error != nil {
		return patient, result.Error
	}

	// Fetch the updated patient record to return
	err := database.DB.First(&patient, patientID).Error
	if err != nil {
		return patient, err
	}

	return patient, nil
}
