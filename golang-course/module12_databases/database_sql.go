package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// ─────── 1. DATABASE/SQL INTERFACE ───────
// Go's `database/sql` package provides a generic interface around SQL databases.
// It DOES NOT include database drivers. You must import the driver for your specific
// database (Postgres, MySQL, SQLite, etc.) using a blank import:
//
// import _ "github.com/lib/pq" 
//
// The blank import calls the driver's `init()` function, which registers the driver
// with `database/sql`. This is a classic Go design pattern.

func demonstrateDatabaseSQL() {
	fmt.Println("\n--- database/sql ---")
	
	// Because we don't have a real database running for this demo,
	// this code focuses on the *patterns* rather than executing them.
	// In a real app, you would use:
	// db, err := sql.Open("postgres", "postgres://user:pass@localhost/mydb?sslmode=disable")
	
	fmt.Println("1. sql.Open() creates a connection POOL, not a single connection.")
	fmt.Println("   It rarely returns an error. It doesn't actually connect to the DB!")
	fmt.Println("   To test the connection, you must call db.Ping().")
	
	fmt.Println("\n2. Executing Queries (No rows returned):")
	fmt.Println("   db.ExecContext(ctx, \"INSERT INTO users (name) VALUES ($1)\", \"Alice\")")
	
	fmt.Println("\n3. Querying Rows (Multiple rows):")
	fmt.Println("   rows, err := db.QueryContext(ctx, \"SELECT id, name FROM users\")")
	fmt.Println("   defer rows.Close() // CRITICAL: Always close rows to release the connection!")
	fmt.Println("   for rows.Next() {")
	fmt.Println("       var id int; var name string")
	fmt.Println("       rows.Scan(&id, &name)")
	fmt.Println("   }")
	
	fmt.Println("\n4. Querying a Single Row:")
	fmt.Println("   var name string")
	fmt.Println("   err := db.QueryRowContext(ctx, \"SELECT name FROM users WHERE id=$1\", 1).Scan(&name)")
	fmt.Println("   if err == sql.ErrNoRows { ... }")
}

// ─────── 2. TRANSACTIONS ───────
// Transactions ensure that a series of operations either all succeed, or all fail.

func simulateTransactionPattern(db *sql.DB) error {
	// 1. Begin the transaction
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	
	// 2. Defer a rollback.
	// If the function returns normally (returning nil error), we will have committed
	// the transaction BEFORE this defer runs. tx.Rollback() does nothing if the
	// transaction is already committed.
	// If the function panics or returns early with an error, the rollback executes!
	defer tx.Rollback()
	
	// 3. Execute queries using the `tx` object, NOT the `db` object.
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - 100 WHERE id = 1")
	if err != nil {
		return err // Rollback will happen
	}
	
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + 100 WHERE id = 2")
	if err != nil {
		return err // Rollback will happen
	}
	
	// 4. Commit if everything succeeded.
	if err := tx.Commit(); err != nil {
		return err
	}
	
	return nil
}
