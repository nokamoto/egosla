package mysql

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nokamoto/egosla/api"
	"gorm.io/gorm"
)

func mockPersistentSubscription(t *testing.T, f func(*PersistentSubscription, sqlmock.Sqlmock)) {
	mockSql(t, func(mock sqlmock.Sqlmock, db *gorm.DB) {
		p := &PersistentSubscription{
			db: db,
		}

		f(p, mock)
	})
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
		t.Run(x.name, func(t *testing.T) {
			mockPersistentSubscription(t, func(p *PersistentSubscription, mock sqlmock.Sqlmock) {
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
