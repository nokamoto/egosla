package mysql

import (
	"errors"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/egosla/api"
	"google.golang.org/protobuf/testing/protocmp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mockPersistentWatcher(t *testing.T, f func(*PersistentWatcher, sqlmock.Sqlmock)) {
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

	f(p, mock)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestPersistentWatcher_Create(t *testing.T) {
	query := func(mock sqlmock.Sqlmock) *sqlmock.ExpectedExec {
		return mock.
			ExpectExec(regexp.QuoteMeta("INSERT INTO `watcher` (`name`,`keywords`) VALUES (?,?)")).
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
			mockPersistentWatcher(t, func(p *PersistentWatcher, mock sqlmock.Sqlmock){
				if x.mock != nil {
					x.mock(mock)
				}

				actual := p.Create(input)

				if !errors.Is(actual, x.expected) {
					t.Errorf("expected %v but actual %v", x.expected, actual)
				}
			})
		})
	}
}


func TestPersistentWatcher_List(t *testing.T){
	offset := 1
	limit := 10

	query := func(mock sqlmock.Sqlmock) *sqlmock.ExpectedQuery {
		return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `watcher` LIMIT 10 OFFSET 1"))
	}

	expected := []*api.Watcher{
		{
			Name: "foo",
			Keywords: []string{"bar", "baz"},
		},
		{
			Name: "qux",
			Keywords: []string{"quux"},
		},
	}

	testcases := []struct {
		name string
		mock func(mock sqlmock.Sqlmock)
		expected []*api.Watcher
		err error
	}{
		{
			name: "ok",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				query(mock).WillReturnRows(
					sqlmock.NewRows([]string{"name", "keywords"}).
					AddRow(expected[0].GetName(), strings.Join(expected[0].GetKeywords(), ",")).
					AddRow(expected[1].GetName(), strings.Join(expected[1].GetKeywords(), ",")),
				)
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

	for _, x := range testcases {
		t.Run(x.name, func(t *testing.T) {
			mockPersistentWatcher(t, func(p *PersistentWatcher, mock sqlmock.Sqlmock){
				if x.mock != nil {
					x.mock(mock)
				}

				actual, err := p.List(offset, limit)

				if !errors.Is(err, x.err) {
					t.Errorf("expected %v but actual %v", x.expected, actual)
				}

				if diff := cmp.Diff(x.expected, actual, protocmp.Transform()); len(diff) != 0 {
					t.Error(diff)
				}
			})
		})
	}
}
