package main

import (
	"testing"
)

func TestGetUserCount(t *testing.T) {
	t.Run("Execute SQL", func(t *testing.T) {
		t.Skip("Skipping DB tests. To run this, you need a mock like DATA-DOG/go-sqlmock or a real DB.")
		/*
			// Example using DATA-DOG/go-sqlmock
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			rows := sqlmock.NewRows([]string{"count"}).AddRow(5)
			mock.ExpectQuery("^SELECT COUNT\\(\\*\\) FROM users$").WillReturnRows(rows)

			count, err := GetUserCount(context.Background(), db)
			if err != nil {
				t.Errorf("error was not expected while getting user count: %s", err)
			}
			if count != 5 {
				t.Errorf("expected count to be 5, got %d", count)
			}
		*/
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Execute SQL", func(t *testing.T) {
		t.Skip("Skipping DB tests.")
		/*
			db, mock, err := sqlmock.New()
			// ... setup mock ...
			mock.ExpectExec("^DELETE FROM users WHERE id = \\$1$").
				WithArgs(10).
				WillReturnResult(sqlmock.NewResult(0, 1)) // 0 insert id, 1 row affected

			affected, err := DeleteUser(context.Background(), db, 10)
			// ... assertions ...
		*/
	})
}
