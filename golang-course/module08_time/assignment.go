package main

import (
	"time"
)

// TODO 1: Implement ParseDate
// It should take a string like "YYYY-MM-DD" and return a time.Time in UTC.
func ParseDate(dateStr string) (time.Time, error) {
	return time.Time{}, nil // Fix me
}

// TODO 2: Implement FormatUSDate
// It should take a time.Time and return a string formatted as "MM/DD/YYYY".
func FormatUSDate(t time.Time) string {
	return "" // Fix me
}

// TODO 3: Implement IsWeekend
// It should return true if the given time is a Saturday or Sunday.
// Hint: look at the time.Weekday() method.
func IsWeekend(t time.Time) bool {
	return false // Fix me
}

// TODO 4: Implement CalculateAge
// It should take a birthdate and a current date, and return the age in full years.
// For example, if born 2000-10-31 and today is 2024-10-30, age is 23.
// If today is 2024-10-31, age is 24.
func CalculateAge(birthdate, today time.Time) int {
	return 0 // Fix me
}
