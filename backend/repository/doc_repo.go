package repository

import (
	"github.com/alyosha-bar/medPortal/models"
	"gorm.io/gorm"
)

type DoctorRepo struct {
	DB *gorm.DB
}

func NewDoctorRepo(db *gorm.DB) *DoctorRepo {
	return &DoctorRepo{DB: db}
}

func (r *DoctorRepo) GetPatientsByDoctor(doctorID uint) ([]models.Patient, error) {
	var patients []models.Patient
	result := r.DB.Where("doctor_id = ?", doctorID).Find(&patients)
	return patients, result.Error
}

func (r *DoctorRepo) UpdateMedicalNotes(doctorID uint, patientID uint, medicalNotes string) (models.Patient, error) {
	var patient models.Patient

	result := r.DB.Model(&patient).
		Where("doctor_id = ? AND id = ?", doctorID, patientID).
		Update("medical_notes", medicalNotes)

	if result.Error != nil {
		return patient, result.Error
	}

	// fetch the updated patient
	err := r.DB.Where("doctor_id = ? AND id = ?", doctorID, patientID).First(&patient).Error

	return patient, err
}
