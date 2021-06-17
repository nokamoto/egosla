package mysql

import (
	"errors"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/fieldmasktest"
	"github.com/nokamoto/egosla/internal/prototest"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"gorm.io/gorm"
)

func mockPersistentWatcher(t *testing.T, name string, f func(*PersistentWatcher, sqlmock.Sqlmock)) {
	t.Run(name, mockSQL(func(mock sqlmock.Sqlmock, db *gorm.DB) {
		p := &PersistentWatcher{
			db: db,
		}

		f(p, mock)
	}))
}

func TestPersistentWatcher_Create(t *testing.T) {
	query := func(mock sqlmock.Sqlmock) *sqlmock.ExpectedExec {
		return mock.
			ExpectExec(regexp.QuoteMeta("INSERT INTO `watcher` (`name`,`keywords`) VALUES (?,?)")).
			WithArgs("foo", "bar,baz")
	}

	input := &api.Watcher{
		Name:     "foo",
		Keywords: []string{"bar", "baz"},
	}

	for _, x := range createMethodTestCases(query) {
		mockPersistentWatcher(t, x.name, func(p *PersistentWatcher, mock sqlmock.Sqlmock) {
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

func TestPersistentWatcher_List(t *testing.T) {
	offset := 1
	limit := 10

	query := func(mock sqlmock.Sqlmock) *sqlmock.ExpectedQuery {
		return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `watcher` LIMIT 10 OFFSET 1"))
	}

	expected := []proto.Message{
		&api.Watcher{
			Name:     "foo",
			Keywords: []string{"bar", "baz"},
		},
		&api.Watcher{
			Name:     "qux",
			Keywords: []string{"quux"},
		},
	}

	rows := sqlmock.NewRows([]string{"name", "keywords"}).AddRow("foo", "bar,baz").AddRow("qux", "quux")

	for _, x := range listMethodTestCases(query, expected, rows) {
		mockPersistentWatcher(t, x.name, func(p *PersistentWatcher, mock sqlmock.Sqlmock) {
			if x.mock != nil {
				x.mock(mock)
			}

			actual, err := p.List(offset, limit)

			var messages []proto.Message
			for _, a := range actual {
				messages = append(messages, a)
			}

			if !errors.Is(err, x.err) {
				t.Errorf("expected %v but actual %v", x.err, err)
			}

			if diff := cmp.Diff(x.expected, messages, protocmp.Transform()); len(diff) != 0 {
				t.Error(diff)
			}
		})
	}
}

func TestPersistentWatcher_Delete(t *testing.T) {
	query := "DELETE FROM `watcher` WHERE name = ?"
	name := "foo"

	for _, x := range deleteMethodTestCases(query, name) {
		mockPersistentWatcher(t, x.name, func(p *PersistentWatcher, mock sqlmock.Sqlmock) {
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

func TestPersistentWatcher_Update(t *testing.T) {
	update := &api.Watcher{
		Name:     "foo",
		Keywords: []string{"bar"},
	}

	updateQuery := func(mock sqlmock.Sqlmock) *sqlmock.ExpectedExec {
		return mock.ExpectExec(regexp.QuoteMeta("UPDATE `watcher` SET `keywords`=? WHERE name = ?"))
	}

	getQuery := func(mock sqlmock.Sqlmock) *sqlmock.ExpectedQuery {
		return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `watcher` WHERE name = ? ORDER BY `watcher`.`name` LIMIT 1"))
	}

	testcases := []struct {
		name       string
		update     *api.Watcher
		updateMask *field_mask.FieldMask
		mock       func(mock sqlmock.Sqlmock)
		expected   *api.Watcher
		err        error
	}{
		{
			name:       "ok",
			update:     update,
			updateMask: fieldmasktest.NewValidFieldMask(t, update, "keywords"),
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()

				updateQuery(mock).
					WithArgs(strings.Join(update.GetKeywords(), ","), update.GetName()).
					WillReturnResult(sqlmock.NewResult(0, 1))

				getQuery(mock).WithArgs(update.GetName()).WillReturnRows(
					sqlmock.NewRows([]string{"name", "keywords"}).
						AddRow(update.GetName(), strings.Join(update.GetKeywords(), ",")),
				)

				mock.ExpectCommit()
			},
			expected: update,
		},
		{
			name:       "err: name field mask",
			update:     update,
			updateMask: fieldmasktest.NewValidFieldMask(t, update, "name"),
			err:        ErrInvalidArgument,
		},
		{
			name:       "err: empty field mask",
			update:     update,
			updateMask: &field_mask.FieldMask{},
			err:        ErrInvalidArgument,
		},
		{
			name:   "err: unknown field mask",
			update: update,
			updateMask: &field_mask.FieldMask{
				Paths: []string{"foo"},
			},
			err: ErrInvalidArgument,
		},
		{
			name:       "err: update unaffected",
			update:     update,
			updateMask: fieldmasktest.NewValidFieldMask(t, update, "keywords"),
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()

				updateQuery(mock).
					WithArgs(strings.Join(update.GetKeywords(), ","), update.GetName()).
					WillReturnResult(sqlmock.NewResult(0, 0))

				mock.ExpectRollback()
			},
			err: ErrUnknown,
		},
		{
			name:       "err: update failed",
			update:     update,
			updateMask: fieldmasktest.NewValidFieldMask(t, update, "keywords"),
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()

				updateQuery(mock).
					WithArgs(strings.Join(update.GetKeywords(), ","), update.GetName()).
					WillReturnError(errors.New("unexpected"))

				mock.ExpectRollback()
			},
			err: ErrUnknown,
		},
		{
			name:       "err: select failed",
			update:     update,
			updateMask: fieldmasktest.NewValidFieldMask(t, update, "keywords"),
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()

				updateQuery(mock).
					WithArgs(strings.Join(update.GetKeywords(), ","), update.GetName()).
					WillReturnResult(sqlmock.NewResult(0, 1))

				getQuery(mock).WithArgs(update.GetName()).WillReturnError(errors.New("unexpected"))

				mock.ExpectRollback()
			},
			err: ErrUnknown,
		},
	}

	for _, x := range testcases {
		mockPersistentWatcher(t, x.name, func(p *PersistentWatcher, mock sqlmock.Sqlmock) {
			if x.mock != nil {
				x.mock(mock)
			}

			actual, err := p.Update(x.update, x.updateMask)

			if diff := cmp.Diff(x.expected, actual, protocmp.Transform()); len(diff) != 0 {
				t.Error(diff)
			}
			if !errors.Is(err, x.err) {
				t.Errorf("expected %v but actual %v", x.err, err)
			}
		})
	}
}

func TestPersistentWatcher_Get(t *testing.T) {
	name := "foo"

	query := "SELECT * FROM `watcher` WHERE name = ? LIMIT 1"

	expected := api.Watcher{
		Name:     "foo",
		Keywords: []string{"bar", "baz"},
	}

	rows := sqlmock.NewRows([]string{"name", "keywords"}).AddRow("foo", "bar,baz")

	empty := sqlmock.NewRows([]string{"name", "keywords"})

	for _, x := range getMethodTestCases(query, name, &expected, rows, empty) {
		mockPersistentWatcher(t, x.name, func(pw *PersistentWatcher, s sqlmock.Sqlmock) {
			if x.mock != nil {
				x.mock(s)
			}

			actual, err := pw.Get(name)

			if err := prototest.Equal(x.expected, actual); err != nil {
				t.Error(err)
			}

			if !errors.Is(err, x.err) {
				t.Errorf("expected %v but actual %v", x.err, err)
			}
		})
	}
}
