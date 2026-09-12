package main

import "fmt"

// ─────── 4. ADVANCED TESTING TECHNIQUES ───────
// We've seen Table-Driven Tests extensively in the assignments.
// What about dependencies like databases or external APIs?

// 4.1 Dependency Injection
// Instead of a function connecting to a database itself, PASS the dependency in.
// This allows you to pass a MOCK dependency during testing.

type Notifier interface {
	Send(message string) error
}

type NotificationService struct {
	client Notifier // Depends on the interface, not a concrete implementation
}

func (s *NotificationService) NotifyUser(msg string) error {
	// ... business logic ...
	return s.client.Send(msg)
}

// In your test:
// type MockNotifier struct { SentMessages []string }
// func (m *MockNotifier) Send(msg string) error { m.SentMessages = append(m.SentMessages, msg); return nil }
//
// func TestNotifyUser(t *testing.T) {
//     mock := &MockNotifier{}
//     svc := NotificationService{client: mock}
//     svc.NotifyUser("hello")
//     if len(mock.SentMessages) != 1 { t.Errorf(...) }
// }

// 4.2 Code Coverage
// Run `go test -cover` to see what percentage of your code is exercised by tests.
// Run `go test -coverprofile=coverage.out` and `go tool cover -html=coverage.out`
// to see a visual, line-by-line breakdown in your browser!

func demonstrateTestingTechniques() {
	fmt.Println("\n--- Testing Techniques ---")
	fmt.Println("See comments in testing_techniques.go for Dependency Injection examples.")
	fmt.Println("Run 'go test -cover' to check code coverage!")
}
