package repository

import (
	"github.com/alyosha-bar/medPortal/models"
	"gorm.io/gorm"
)

type ReceptionRepo struct {
	DB *gorm.DB
}

func NewReceptionRepo(db *gorm.DB) *ReceptionRepo {
	return &ReceptionRepo{DB: db}
}

func (r *ReceptionRepo) GetAllPatients() ([]models.Patient, error) {
	var patients []models.Patient
	result := r.DB.Find(&patients)
	return patients, result.Error
}

func (r *ReceptionRepo) GetPatient(patientID uint) (models.Patient, error) {
	var patient models.Patient
	result := r.DB.Where("id = ?", patientID).First(&patient)
	return patient, result.Error
}

func (r *ReceptionRepo) RegisterPatient(patient models.Patient) (models.Patient, error) {
	result := r.DB.Create(&patient)
	return patient, result.Error
}

func (r *ReceptionRepo) DeletePatientProfile(patientID uint) error {
	result := r.DB.Delete(&models.Patient{}, patientID)
	return result.Error
}

func (r *ReceptionRepo) UpdateField(patientID uint, field string, value interface{}) (models.Patient, error) {
	var patient models.Patient

	// perform update
	if err := r.DB.Model(&models.Patient{}).
		Where("id = ?", patientID).
		Update(field, value).Error; err != nil {
		return patient, err
	}

	// return updated patient
	err := r.DB.First(&patient, patientID).Error
	return patient, err
}

func (r *ReceptionRepo) GetAllDoctors() ([]models.User, error) {
	var doctors []models.User
	result := r.DB.Select("id, username").Where("role = ?", "doctor").Find(&doctors)
	return doctors, result.Error
}

func (r *ReceptionRepo) AssignPatient(patientID uint, doctorID uint) (models.Patient, error) {
	var patient models.Patient

	result := r.DB.Model(&patient).
		Where("id = ?", patientID).
		Update("doctor_id", doctorID)

	if result.Error != nil {
		return patient, result.Error
	}

	err := r.DB.First(&patient, patientID).Error
	if err != nil {
		return patient, err
	}

	return patient, nil
}
