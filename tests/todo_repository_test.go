package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"todolist/repository"
)

func TestGetTodosByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init database mock %s", err)
	}
	defer db.Close()

	repo := repository.NewRepo(db, nil)

	rows := sqlmock.NewRows([]string{"id", "title", "completed"}).
		AddRow(1, "test 1 ", false).
		AddRow(2, "test 2", true)

	mock.ExpectQuery("SELECT id, title, completed FROM todos WHERE user_id = ?").
		WithArgs(1).
		WillReturnRows(rows)

	todos, err := repo.GetTodos(1)
	assert.NoError(t, err)
	assert.Len(t, todos, 2)
	assert.Equal(t, "test 1", todos[0].Title)
}
