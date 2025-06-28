package services

import (
	"github.com/alyosha-bar/medPortal/models"
	"github.com/alyosha-bar/medPortal/repository"
)

func GetPatientsByDoctor(doctorID uint) ([]models.Patient, error) {
	return repository.GetPatientsByDoctor(doctorID)
}
