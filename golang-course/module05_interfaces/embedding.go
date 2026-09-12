package main

import "fmt"

// ─────── 5. COMPOSITION OVER INHERITANCE (EMBEDDING) ───────
// Go does not have `class`, `extends`, or inheritance.
// Instead, Go uses Struct Embedding to achieve composition.
//
// "Prefer composition over inheritance" is a classic OOP design principle.
// Go enforces this at the language level.

type User struct {
	ID   int
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("User[%d]: %s", u.ID, u.Name)
}

// Admin embeds the User struct.
// It is NOT a subclass of User. It HAS a User.
// However, the fields and methods of the embedded User are "promoted" to the Admin level
// for syntactic convenience.
type Admin struct {
	User        // Anonymous embedded field
	Permissions []string
}

func demonstrateEmbedding() {
	fmt.Println("\n--- Struct Embedding (Composition) ---")
	
	admin := Admin{
		User: User{
			ID:   1,
			Name: "Alice",
		},
		Permissions: []string{"write", "delete"},
	}

	// We can access Name directly because it is promoted.
	fmt.Printf("Admin Name: %s\n", admin.Name) 
	
	// We can also access it explicitly via the embedded type name.
	fmt.Printf("Admin User Name: %s\n", admin.User.Name)

	// Since User has a String() method, does Admin?
	// YES! Methods are also promoted.
	fmt.Printf("Admin Stringer: %s\n", admin.String())
	
	// IMPORTANT GOTCHA:
	// Even though Admin has all of User's methods, an Admin is NOT a User.
	// You cannot pass an Admin to a function that expects a User.
	// (Unless you pass admin.User).
	// To achieve polymorphism, you MUST use interfaces, not embedding.
}
