package mysql

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/prototest"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newStdPersistent(t *testing.T, f func(sqlmock.Sqlmock)) *StdPersistent {
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

	if f != nil {
		f(mock)
	}

	return &StdPersistent{
		std:   &std{db: db},
		model: &WatcherModel{},
	}
}

func newRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"name", "keywords"})
}

func TestStdPersistent_Create(t *testing.T) {
	requested := &api.Watcher{
		Name:     "foo",
		Keywords: []string{"bar"},
	}

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
			sp := newStdPersistent(t, testcase.mock)
			if err := sp.Create(requested); !errors.Is(err, testcase.expected) {
				t.Errorf("expected %v but actual %v", testcase.expected, err)
			}
		})
	}
}

func TestStdPersistent_Get(t *testing.T) {
	got := &api.Watcher{
		Name:     "foo",
		Keywords: []string{"bar"},
	}

	query := func(s sqlmock.Sqlmock) *sqlmock.ExpectedQuery {
		return s.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `watcher` WHERE name = ? LIMIT 1")).WithArgs(got.GetName())
	}

	testcases := []struct {
		name     string
		mock     func(sqlmock.Sqlmock)
		expected proto.Message
		err      error
	}{
		{
			name: "ok",
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				query(s).WillReturnRows(
					newRows().AddRow(got.GetName(), strings.Join(got.GetKeywords(), ",")),
				)
				s.ExpectCommit()
			},
			expected: got,
		},
		{
			name: "not found",
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				query(s).WillReturnRows(newRows())
				s.ExpectRollback()
			},
			err: ErrNotFound,
		},
		{
			name: "unexpected error",
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				query(s).WillReturnError(errors.New("unexpected"))
				s.ExpectRollback()
			},
			err: ErrUnknown,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			sp := newStdPersistent(t, testcase.mock)
			actual, err := sp.Get(got.GetName())
			if err := prototest.Equal(testcase.expected, actual); err != nil {
				t.Error(err)
			}
			if !errors.Is(err, testcase.err) {
				t.Errorf("expected %v but actual %v", testcase.err, err)
			}
		})
	}
}

func TestStdPersistent_List(t *testing.T) {
	got := []*api.Watcher{
		{
			Name: "foo",
		},
		{
			Name: "bar",
		},
	}

	offset, limit := 1, 2

	query := func(s sqlmock.Sqlmock) *sqlmock.ExpectedQuery {
		return s.ExpectQuery(regexp.QuoteMeta(fmt.Sprintf("SELECT * FROM `watcher` LIMIT %d OFFSET %d", limit, offset)))
	}

	testcases := []struct {
		name     string
		mock     func(sqlmock.Sqlmock)
		expected []proto.Message
		err      error
	}{
		{
			name: "ok: empty",
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				query(s).WillReturnRows(newRows())
				s.ExpectCommit()
			},
		},
		{
			name: "ok",
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				query(s).WillReturnRows(newRows().AddRow(got[0].GetName(), "").AddRow(got[1].GetName(), ""))
				s.ExpectCommit()
			},
			expected: []proto.Message{got[0], got[1]},
		},
		{
			name: "unexpected error",
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				query(s).WillReturnError(errors.New("unexpected"))
				s.ExpectRollback()
			},
			err: ErrUnknown,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			sp := newStdPersistent(t, testcase.mock)
			actual, err := sp.List(offset, limit)
			if err := prototest.EqualSlice(testcase.expected, actual); err != nil {
				t.Error(err)
			}
			if !errors.Is(err, testcase.err) {
				t.Errorf("expected %v but actual %v", testcase.err, err)
			}
		})
	}
}
