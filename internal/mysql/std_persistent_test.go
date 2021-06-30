package mysql

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nokamoto/egosla/api"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	gdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

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
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}

		gdb.Close()
	})

	return db, mock
}

func TestStdPersistent_Create(t *testing.T) {
	created := &watcher{
		Name:     "foo",
		Keywords: "bar",
	}

	query := func(s sqlmock.Sqlmock) *sqlmock.ExpectedExec {
		return s.ExpectExec(regexp.QuoteMeta("INSERT INTO `watcher` (`name`,`keywords`) VALUES (?,?)")).WithArgs(created.Name, created.Keywords)
	}

	testcases := []struct {
		name     string
		mock     func(sqlmock.Sqlmock)
		expected error
	}{
		{
			name: "ok",
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				query(s).WillReturnResult(sqlmock.NewResult(0, 1))
				s.ExpectCommit()
			},
		},
		{
			name: "unexpected error",
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				query(s).WillReturnError(errors.New("unexpected"))
				s.ExpectRollback()
			},
			expected: ErrUnknown,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			db, mock := newTestDB(t)

			if testcase.mock != nil {
				testcase.mock(mock)
			}

			sp := &StdPersistent{
				std: &std{db: db},
				convert: func(_ proto.Message) (interface{}, error) {
					return created, nil
				},
			}

			dummy := &api.Watcher{}

			if err := sp.Create(dummy); !errors.Is(err, testcase.expected) {
				t.Errorf("expected %v but actual %v", testcase.expected, err)
			}
		})
	}
}
