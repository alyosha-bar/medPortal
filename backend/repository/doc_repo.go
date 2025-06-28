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

func UpdateMedicalNotes(doctorID uint, patientID uint, medicalNotes string) (models.Patient, error) {
	var patient models.Patient

	result := database.DB.Model(&patient).
		Where("doctor_id = ? AND id = ?", doctorID, patientID).
		Update("medical_notes", medicalNotes)

	if result.Error != nil {
		return patient, result.Error
	}

	// fetch the updated patient
	err := database.DB.Where("doctor_id = ? AND id = ?", doctorID, patientID).First(&patient).Error

	return patient, err
}
