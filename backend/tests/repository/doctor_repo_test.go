package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alyosha-bar/medPortal/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gdb, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	return gdb, mock, func() {
		db.Close()
	}

}

func TestGetPatientByDoctor(t *testing.T) {
	db, mock, close := setupMockDB(t)
	defer close()

	repo := repository.NewDoctorRepo(db)

	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "age", "doctor_id", "medical_notes"}).
		AddRow(1, "John", "Doe", 30, 5, "Healthy").
		AddRow(2, "Jane", "Roe", 25, 5, "Needs follow-up")

	mock.ExpectQuery(`SELECT \* FROM "patients" WHERE doctor_id = \$1`).
		WithArgs(5).
		WillReturnRows(rows)

	patients, err := repo.GetPatientsByDoctor(5)

	assert.NoError(t, err)
	assert.Len(t, patients, 2)
	assert.Equal(t, "John", patients[0].Firstname)
	assert.Equal(t, "Healthy", patients[0].MedicalNotes)
	assert.Equal(t, uint(5), *patients[0].DoctorID)
}

func TestUpdateMedicalNotes(t *testing.T) {
	db, mock, close := setupMockDB(t)
	defer close()

	repo := repository.NewDoctorRepo(db)

	// execute update
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "patients" SET "medical_notes"=\$1 WHERE doctor_id = \$2 AND id = \$3`).
		WithArgs("Updated notes", 5, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// get the update above
	rows := sqlmock.NewRows([]string{"id", "name", "doctor_id", "medical_notes"}).
		AddRow(1, "John Doe", 5, "Updated notes")

	mock.ExpectQuery(`SELECT \* FROM "patients" WHERE doctor_id = \$1 AND id = \$2 ORDER BY "patients"\."id" LIMIT \$3`).
		WithArgs(5, 1, 1).
		WillReturnRows(rows)

	patient, err := repo.UpdateMedicalNotes(5, 1, "Updated notes")

	assert.NoError(t, err)
	assert.Equal(t, "Updated notes", patient.MedicalNotes)
}
