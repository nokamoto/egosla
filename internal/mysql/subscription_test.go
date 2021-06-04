package mysql

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/egosla/api"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"gorm.io/gorm"
)

func mockPersistentSubscription(t *testing.T, name string, f func(*PersistentSubscription, sqlmock.Sqlmock)) {
	t.Run(name, mockSQL(func(mock sqlmock.Sqlmock, db *gorm.DB) {
		p := &PersistentSubscription{
			db: db,
		}

		f(p, mock)
	}))
}

func TestPersistentSubscription_Create(t *testing.T) {
	query := func(mock sqlmock.Sqlmock) *sqlmock.ExpectedExec {
		return mock.
			ExpectExec(regexp.QuoteMeta("INSERT INTO `subscription` (`name`,`watcher`) VALUES (?,?)")).
			WithArgs("foo", "bar")
	}

	input := &api.Subscription{
		Name:    "foo",
		Watcher: "bar",
	}

	for _, x := range createMethodTestCases(query) {
		mockPersistentSubscription(t, x.name, func(p *PersistentSubscription, mock sqlmock.Sqlmock) {
			if x.mock != nil {
				x.mock(mock)
			}

			actual := p.Create(input)

			if !errors.Is(actual, x.expected) {
				t.Errorf("expected %v but actual %v", x.expected, actual)
			}
		})
	}
}

func TestPersistentSubscription_List(t *testing.T) {
	offset := 100
	limit := 200

	query := func(mock sqlmock.Sqlmock) *sqlmock.ExpectedQuery {
		return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `subscription` LIMIT 200 OFFSET 100"))
	}

	expected := []proto.Message{
		&api.Subscription{
			Name:    "foo",
			Watcher: "bar",
		},
		&api.Subscription{
			Name:    "baz",
			Watcher: "qux",
		},
	}

	rows := sqlmock.NewRows([]string{"name", "watcher"}).AddRow("foo", "bar").AddRow("baz", "qux")

	for _, x := range listMethodTestCases(query, expected, rows) {
		mockPersistentSubscription(t, x.name, func(p *PersistentSubscription, mock sqlmock.Sqlmock) {
			if x.mock != nil {
				x.mock(mock)
			}

			actual, err := p.List(offset, limit)

			var messages []proto.Message
			for _, a := range actual {
				messages = append(messages, a)
			}

			if diff := cmp.Diff(x.expected, messages, protocmp.Transform()); len(diff) != 0 {
				t.Error(diff)
			}

			if !errors.Is(err, x.err) {
				t.Errorf("expected %v but actual %v", x.expected, actual)
			}
		})
	}
}

func TestPersistentSubscription_Delete(t *testing.T) {
	query := "DELETE FROM `subscription` WHERE name = ?"
	name := "foo"

	for _, x := range deleteMethodTestCases(query, name) {
		mockPersistentSubscription(t, x.name, func(p *PersistentSubscription, mock sqlmock.Sqlmock) {
			if x.mock != nil {
				x.mock(mock)
			}

			actual := p.Delete(name)

			if !errors.Is(actual, x.expected) {
				t.Errorf("expected %v but actual %v", x.expected, actual)
			}
		})
	}
}
