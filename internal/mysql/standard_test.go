package mysql

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mockSql(t *testing.T, f func(sqlmock.Sqlmock, *gorm.DB)) {
	gdb, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer gdb.Close()

	db, err := gorm.Open(
		mysql.Dialector{
			Config: &mysql.Config{
				DriverName:                "mysql",
				Conn:                      gdb,
				SkipInitializeWithVersion: true,
			},
		},
		&gorm.Config{},
	)
	if err != nil {
		t.Error(err)
	}

	f(mock, db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func createMethodTestCases(query func(sqlmock.Sqlmock) *sqlmock.ExpectedExec) []struct {
	name     string
	mock     func(mock sqlmock.Sqlmock)
	expected error
} {
	return []struct {
		name     string
		mock     func(mock sqlmock.Sqlmock)
		expected error
	}{
		{
			name: "ok",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				query(mock).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expected: nil,
		},
		{
			name: "unexpected error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				query(mock).WillReturnError(errors.New("unexpected"))
				mock.ExpectRollback()
			},
			expected: ErrUnknown,
		},
	}
}
