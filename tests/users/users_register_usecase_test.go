package main

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Age      int
}

func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	t.Run("add user successfully", func(t *testing.T) {
		// Expectation: mock the count query to return 0 for the specified email
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
			WithArgs("john.doe@example.com").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		// Expectation: mock the transaction beginning
		mock.ExpectBegin()

		// Expectation: mock the INSERT query to return the user ID
		mock.ExpectExec("^INSERT INTO \"users\" (.+)$").
			WillReturnResult(sqlmock.NewResult(1, 1)) // Assuming user ID 1 is returned

		// Expectation: mock the transaction commit
		mock.ExpectCommit()

		// Call the function under test
		err := AddUser(gormDB, "John Doe", "john.doe@example.com", 30)

		// Verify no error occurred
		assert.NoError(t, err)

		// Verify that all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func AddUser(db *gorm.DB, fullname, email string, age int) error {
	user := User{Fullname: fullname, Email: email, Age: age}

	// Check if email already exists
	var count int64
	db.Model(&User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return errors.New("email already exists")
	}

	// Save the new user
	result := db.Create(&user)
	return result.Error
}
