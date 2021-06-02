package mysql

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mockSql(f func(sqlmock.Sqlmock, *gorm.DB)) func(*testing.T) {
	return func(t *testing.T) {
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

func listMethodTestCases(query func(sqlmock.Sqlmock) *sqlmock.ExpectedQuery, expected []proto.Message, rows *sqlmock.Rows) []struct {
	name     string
	mock     func(mock sqlmock.Sqlmock)
	expected []proto.Message
	err      error
} {
	return []struct {
		name     string
		mock     func(mock sqlmock.Sqlmock)
		expected []proto.Message
		err      error
	}{
		{
			name: "ok",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				query(mock).WillReturnRows(rows)
				mock.ExpectCommit()
			},
			expected: expected,
		},
		{
			name: "unexpected error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				query(mock).WillReturnError(errors.New("unexpected"))
				mock.ExpectRollback()
			},
			err: ErrUnknown,
		},
	}
}

func deleteMethodTestCases(query string, name string) []struct {
	name     string
	mock     func(mock sqlmock.Sqlmock)
	expected error
} {
	q := func(mock sqlmock.Sqlmock) *sqlmock.ExpectedExec {
		return mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(name)
	}

	return []struct {
		name     string
		mock     func(mock sqlmock.Sqlmock)
		expected error
	}{
		{
			name: "ok",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				q(mock).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "ok: undeleted",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				q(mock).WillReturnResult(sqlmock.NewResult(0, 0))
				mock.ExpectCommit()
			},
		},
		{
			name: "unexpected error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				q(mock).WillReturnError(errors.New("unexpected"))
				mock.ExpectRollback()
			},
			expected: ErrUnknown,
		},
	}
}
