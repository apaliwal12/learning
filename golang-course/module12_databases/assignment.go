package main

import (
	"context"
	"database/sql"
)

// TODO 1: Implement GetUserCount
// It should execute a query "SELECT COUNT(*) FROM users" on the provided db.
// Return the count and any error.
// Use QueryRowContext.
func GetUserCount(ctx context.Context, db *sql.DB) (int, error) {
	return 0, nil // Fix me
}

// TODO 2: Implement DeleteUser
// It should execute "DELETE FROM users WHERE id = $1" with the given ID.
// It should return the number of rows affected.
// Hint: db.ExecContext returns a sql.Result, which has a RowsAffected() method.
func DeleteUser(ctx context.Context, db *sql.DB, id int) (int64, error) {
	return 0, nil // Fix me
}

// TODO 3: Define a User struct
// It should have ID (int), Name (string), and Email (string).
// type User struct { ... }

// TODO 4: Implement GetAllUsers
// It should execute "SELECT id, name, email FROM users"
// It must iterate over the rows, scan them into User structs, and return a slice of Users.
// Don't forget to close the rows and check for rows.Err()!
// func GetAllUsers(ctx context.Context, db *sql.DB) ([]User, error) { ... }
