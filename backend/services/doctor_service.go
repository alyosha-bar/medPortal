package services

import (
	"github.com/alyosha-bar/medPortal/models"
	"github.com/alyosha-bar/medPortal/repository"
)

type DoctorService struct {
	Repo *repository.DoctorRepo
}

func NewDoctorService(repo *repository.DoctorRepo) *DoctorService {
	return &DoctorService{Repo: repo}
}

func (s *DoctorService) GetPatientsByDoctor(doctorID uint) ([]models.Patient, error) {
	return s.Repo.GetPatientsByDoctor(doctorID)
}

func (s *DoctorService) UpdateMedicalNotes(doctorID uint, patientID uint, medicalNotes string) (models.Patient, error) {
	return s.Repo.UpdateMedicalNotes(doctorID, patientID, medicalNotes)
}
