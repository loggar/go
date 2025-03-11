package sqlmock_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// Function that runs a NamedExecContext
func updateUser(ctx context.Context, db *sql.DB, id int, name string) error {
	query := `UPDATE users SET name = :name WHERE id = :id`
	_, err := db.ExecContext(ctx, query, sql.Named("name", name), sql.Named("id", id))
	return err
}

// Test function that tests the updateUser function using ExpectExec and regexp.QuoteMeta
func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.Background()
	id := 1
	name := "John Doe"

	query := `UPDATE users SET name = :name WHERE id = :id`
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(sql.Named("name", name), sql.Named("id", id)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = updateUser(ctx, db, id, name)
	if err != nil {
		t.Errorf("error was not expected while updating user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
