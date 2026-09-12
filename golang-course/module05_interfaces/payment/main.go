package main

import (
	"fmt"
	"strings"
)

// ─────── PAYMENT PROCESSOR PROJECT ───────
// This project demonstrates how interfaces allow us to build a flexible
// payment processing system where different payment methods can be plugged
// in without changing the core business logic.

// 1. Define the Interface
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	GetMethodName() string
}

// 2. Implement Credit Card Processor
type CreditCard struct {
	CardNumber string
	CardHolder string
}

func (cc CreditCard) ProcessPayment(amount float64) error {
	// In a real app, we'd talk to Stripe or Visa here
	fmt.Printf("[CreditCard] Charging $%.2f to card ending in %s\n", 
		amount, cc.CardNumber[len(cc.CardNumber)-4:])
	return nil
}

func (cc CreditCard) GetMethodName() string {
	return "Credit Card"
}

// 3. Implement PayPal Processor
type PayPal struct {
	Email string
}

func (pp PayPal) ProcessPayment(amount float64) error {
	fmt.Printf("[PayPal] Transferring $%.2f from account %s\n", amount, pp.Email)
	return nil
}

func (pp PayPal) GetMethodName() string {
	return "PayPal"
}

// 4. Core Business Logic (Depends on Interface, NOT concretions)
func Checkout(amount float64, processor PaymentProcessor) {
	fmt.Printf("Starting checkout of $%.2f via %s...\n", amount, processor.GetMethodName())
	
	err := processor.ProcessPayment(amount)
	if err != nil {
		fmt.Printf("Checkout failed: %v\n", err)
		return
	}
	
	fmt.Println("Checkout successful! Thank you for your purchase.")
	fmt.Println(strings.Repeat("-", 40))
}

// 5. Mock Processor for Unit Testing
// We can easily create a mock because Checkout only depends on the interface!
type MockProcessor struct {
	ShouldFail bool
}

func (m MockProcessor) ProcessPayment(amount float64) error {
	if m.ShouldFail {
		return fmt.Errorf("mock processor forced failure")
	}
	fmt.Printf("[Mock] Successfully processed $%.2f\n", amount)
	return nil
}

func (m MockProcessor) GetMethodName() string {
	return "Mock Processor"
}

func main() {
	fmt.Println("--- Payment Processing System ---")

	cc := CreditCard{CardNumber: "1234567890123456", CardHolder: "Alice"}
	pp := PayPal{Email: "bob@example.com"}

	// Use Credit Card
	Checkout(99.99, cc)

	// Use PayPal
	Checkout(45.50, pp)

	// Use Mock (Test environment)
	fmt.Println("Running test suite...")
	successMock := MockProcessor{ShouldFail: false}
	Checkout(10.00, successMock)
	
	failMock := MockProcessor{ShouldFail: true}
	Checkout(10.00, failMock)
}
