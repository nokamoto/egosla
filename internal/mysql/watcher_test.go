package mysql

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nokamoto/egosla/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestPersistentWatcher_Create(t *testing.T) {
	query := func(mock sqlmock.Sqlmock) *sqlmock.ExpectedExec {
		return mock.
			ExpectExec(regexp.QuoteMeta("INSERT INTO `watchers` (`name`,`keywords`) VALUES (?,?)")).
			WithArgs("foo", "bar,baz")
	}

	input := &api.Watcher {
		Name: "foo",
		Keywords: []string{"bar", "baz"},
	}

	testcases := []struct {
		name string
		mock func(mock sqlmock.Sqlmock)
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

	for _, x := range testcases {
		t.Run(x.name, func(t *testing.T) {
			gdb, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer gdb.Close()

			db, err := gorm.Open(
				mysql.Dialector{
					Config: &mysql.Config{
						DriverName: "mysql", 
						Conn: gdb, 
						SkipInitializeWithVersion: true,
					},
				}, 
				&gorm.Config{},
			)
			if err != nil {
				t.Error(err)
			}

			p := &PersistentWatcher{
				db: db,
			}

			if x.mock != nil {
				x.mock(mock)
			}

			actual := p.Create(input)

			if !errors.Is(actual, x.expected) {
				t.Errorf("expected %v but actual %v", x.expected, actual)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
		})
	}
}
