package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alyosha-bar/medPortal/models"
	"github.com/alyosha-bar/medPortal/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPatients(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "age", "gender", "doctor_id", "medical_notes"}).
		AddRow(1, "Alice", "Smith", 30, "Female", nil, "Note A").
		AddRow(2, "Bob", "Brown", 40, "Male", nil, "Note B")

	mock.ExpectQuery(`SELECT \* FROM "patients"`).WillReturnRows(rows)

	repo := repository.NewReceptionRepo(db)
	patients, err := repo.GetAllPatients()

	assert.NoError(t, err)
	assert.Len(t, patients, 2)
	assert.Equal(t, "Alice", patients[0].Firstname)
}

func TestGetPatient(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "age", "gender", "doctor_id", "medical_notes"}).
		AddRow(1, "Charlie", "Lee", 28, "Nonbinary", nil, "Note C")

	mock.ExpectQuery(`SELECT \* FROM "patients" WHERE id = \$1 ORDER BY "patients"\."id" LIMIT \$2`).
		WithArgs(1, 1).WillReturnRows(rows)

	repo := repository.NewReceptionRepo(db)
	patient, err := repo.GetPatient(1)

	assert.NoError(t, err)
	assert.Equal(t, "Charlie", patient.Firstname)
}

func TestRegisterPatient(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	patient := models.Patient{Firstname: "Daisy", Lastname: "Hopper", Age: 33, Gender: "Female"}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "patients"`).
		WithArgs(patient.Firstname, patient.Lastname, patient.Age, patient.Gender, sqlmock.AnyArg(), "").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	repo := repository.NewReceptionRepo(db)
	_, err := repo.RegisterPatient(patient)

	assert.NoError(t, err)
}

func TestDeletePatientProfile(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "patients" WHERE "patients"\."id" = \$1`).WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := repository.NewReceptionRepo(db)
	err := repo.DeletePatientProfile(1)

	assert.NoError(t, err)
}

func TestUpdateField(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "patients" SET "age"=\$1 WHERE id = \$2`).
		WithArgs(35, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "age", "gender", "doctor_id", "medical_notes"}).
		AddRow(1, "Elena", "Joy", 35, "Female", nil, "Updated")

	mock.ExpectQuery(`SELECT \* FROM "patients" WHERE "patients"\."id" = \$1 ORDER BY "patients"\."id" LIMIT \$2`).
		WithArgs(1, 1).WillReturnRows(rows)

	repo := repository.NewReceptionRepo(db)
	patient, err := repo.UpdateField(1, "age", 35)

	assert.NoError(t, err)
	assert.Equal(t, uint(35), patient.Age)
}

func TestGetAllDoctors(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"id", "username"}).
		AddRow(1, "dr_who").
		AddRow(2, "dr_strange")

	mock.ExpectQuery(`SELECT id, username FROM "users" WHERE role = \$1`).
		WithArgs("doctor").WillReturnRows(rows)

	repo := repository.NewReceptionRepo(db)
	doctors, err := repo.GetAllDoctors()

	assert.NoError(t, err)
	assert.Len(t, doctors, 2)
	assert.Equal(t, "dr_who", doctors[0].Username)
}

func TestAssignPatient(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "patients" SET "doctor_id"=\$1 WHERE id = \$2`).
		WithArgs(2, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "age", "gender", "doctor_id", "medical_notes"}).
		AddRow(1, "Finn", "Miller", 27, "Male", 2, "Reassigned")

	mock.ExpectQuery(`SELECT \* FROM "patients" WHERE "patients"\."id" = \$1 ORDER BY "patients"\."id" LIMIT \$2`).
		WithArgs(1, 1).WillReturnRows(rows)

	repo := repository.NewReceptionRepo(db)
	patient, err := repo.AssignPatient(1, 2)

	assert.NoError(t, err)
	assert.Equal(t, uint(2), *patient.DoctorID)
}
