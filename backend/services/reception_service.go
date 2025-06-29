package services

import (
	"github.com/alyosha-bar/medPortal/models"
	"github.com/alyosha-bar/medPortal/repository"
)

type ReceptionService struct {
	Repo *repository.ReceptionRepo
}

func NewReceptionService(repo *repository.ReceptionRepo) *ReceptionService {
	return &ReceptionService{Repo: repo}
}

func (s *ReceptionService) GetAllPatients() ([]models.Patient, error) {
	return s.Repo.GetAllPatients()
}

func (s *ReceptionService) GetPatient(patientID uint) (models.Patient, error) {
	return s.Repo.GetPatient(patientID)
}

func (s *ReceptionService) RegisterPatient(patient models.Patient) (models.Patient, error) {
	return s.Repo.RegisterPatient(patient)
}

func (s *ReceptionService) DeletePatientProfile(patientID uint) error {
	return s.Repo.DeletePatientProfile(patientID)
}

func (s *ReceptionService) UpdateField(patientID uint, field string, value interface{}) (models.Patient, error) {
	return s.Repo.UpdateField(patientID, field, value)
}

func (s *ReceptionService) GetAllDoctors() ([]models.User, error) {
	return s.Repo.GetAllDoctors()
}

func (s *ReceptionService) AssignPatient(patientID uint, doctorID uint) (models.Patient, error) {
	return s.Repo.AssignPatient(patientID, doctorID)
}
