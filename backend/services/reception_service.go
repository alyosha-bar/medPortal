package services

import (
	"github.com/alyosha-bar/medPortal/models"
	"github.com/alyosha-bar/medPortal/repository"
)

func GetAllPatients() ([]models.Patient, error) {
	return repository.GetAllPatients()
}

func GetPatient(patientID uint) (models.Patient, error) {
	return repository.GetPatient(patientID)
}

func RegisterPatient(patient models.Patient) (models.Patient, error) {
	return repository.RegisterPatient(patient)
}

func DeletePatientProfile(patientID uint) error {
	return repository.DeletePatientProfile(patientID)
}

func GetAllDoctors() ([]models.User, error) {
	return repository.GetAllDoctors()
}

func AssignPatient(patientID uint, doctorID uint) (models.Patient, error) {
	return repository.AssignPatient(patientID, doctorID)
}
