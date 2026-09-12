package main

import (
	"fmt"
)

// ─────── 1. IMPLICIT INTERFACES ───────
// In Go, interfaces are satisfied IMPLICITLY.
// There is no `implements` keyword. If a type has the methods described by an
// interface, it implements the interface.
//
// Rhetoric: Why implicit? 
// 1. It decouples packages. You can define an interface in your package that
//    is satisfied by a type in another package, without modifying the other package.
// 2. It encourages small, single-method interfaces (like io.Reader, io.Writer).

// Notifier defines behavior for sending notifications.
type Notifier interface {
	Notify(message string) error
}

// EmailService implements Notifier (implicitly)
type EmailService struct {
	Email string
}

func (e EmailService) Notify(msg string) error {
	fmt.Printf("Sending email to %s: %s\n", e.Email, msg)
	return nil
}

// SMSService implements Notifier (implicitly)
type SMSService struct {
	PhoneNumber string
}

func (s SMSService) Notify(msg string) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, msg)
	return nil
}

// SendAlert accepts ANY type that implements the Notifier interface.
// This is the core of polymorphism in Go.
func SendAlert(n Notifier, msg string) {
	err := n.Notify(msg)
	if err != nil {
		fmt.Printf("Failed to send alert: %v\n", err)
	}
}

func demonstrateInterfaces() {
	fmt.Println("\n--- Interfaces & Polymorphism ---")
	
	email := EmailService{Email: "admin@example.com"}
	sms := SMSService{PhoneNumber: "555-0199"}

	SendAlert(email, "System is down!")
	SendAlert(sms, "System is down!")
}

// ─────── 2. THE EMPTY INTERFACE (any) ───────
// `interface{}` (or the newer alias `any` introduced in Go 1.18) specifies zero methods.
// Since every type implements at least zero methods, `any` can hold a value of ANY type.

func PrintAnything(v any) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

func demonstrateEmptyInterface() {
	fmt.Println("\n--- The Empty Interface (any) ---")
	
	PrintAnything(42)
	PrintAnything("hello")
	PrintAnything(EmailService{})
}

// ─────── 3. TYPE ASSERTIONS & SWITCHES ───────
// To get the concrete value out of an interface, we use type assertions.

func demonstrateTypeAssertions() {
	fmt.Println("\n--- Type Assertions & Switches ---")
	
	var i any = "hello"

	// Type Assertion (Panic if wrong)
	// s := i.(string) 

	// Safe Type Assertion (comma ok idiom)
	s, ok := i.(string)
	if ok {
		fmt.Printf("It's a string: %s\n", s)
	}

	// Type Switch
	switch v := i.(type) {
	case string:
		fmt.Printf("String length: %d\n", len(v))
	case int:
		fmt.Printf("Double the int: %d\n", v*2)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// ─────── 4. THE NIL INTERFACE GOTCHA ───────
// Internally, an interface is a tuple of (Type, Value).
// An interface is ONLY nil if BOTH Type and Value are nil.
// If you assign a nil pointer of a concrete type to an interface, the interface is NOT nil!

func demonstrateNilInterface() {
	fmt.Println("\n--- The Nil Interface Gotcha ---")
	
	var n Notifier
	fmt.Printf("Is n nil? %v (Type: %T, Value: %v)\n", n == nil, n, n) // true

	// Create a typed pointer, set it to nil
	var emailPtr *EmailService // nil pointer
	
	// Assign it to the interface
	n = emailPtr 
	
	// The interface now holds (Type: *EmailService, Value: nil)
	// Therefore, the interface itself is NOT nil!
	fmt.Printf("Is n nil? %v (Type: %T, Value: %v)\n", n == nil, n, n) // FALSE!
	
	// This is a very common source of bugs in Go when returning errors.
	// ALWAYS return explicit `nil`, not a nil pointer variable, when returning an error.
}
