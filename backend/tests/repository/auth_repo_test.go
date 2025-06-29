package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alyosha-bar/medPortal/models"
	"github.com/alyosha-bar/medPortal/repository"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	user := models.User{
		Username: "newuser",
		Password: "hashedpass",
		Role:     "reception",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WithArgs(user.Username, user.Password, user.Role).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	repo := repository.NewAuthRepo(db)
	err := repo.SignUp(user)
	assert.NoError(t, err)
}

func TestGetUserByUsername(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE username = \$1 ORDER BY "users"\."id" LIMIT \$2`).
		WithArgs("admin", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "role"}).
			AddRow(1, "admin", "hashedpass", "admin"))

	repo := repository.NewAuthRepo(db)
	user, err := repo.GetUserByUsername("admin")

	assert.NoError(t, err)
	assert.Equal(t, "admin", user.Username)
	assert.Equal(t, "admin", user.Role)
}
