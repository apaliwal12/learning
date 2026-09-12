package main

// TODO 1: Implement ReverseString
// It should return a new string with the characters reversed.
// WARNING: Remember that strings can contain multi-byte UTF-8 characters (runes).
// Reversing the bytes will corrupt the string! You must reverse the RUNES.
func ReverseString(s string) string {
	return "" // Fix me
}

// TODO 2: Implement IsPalindrome
// It should return true if the string reads the same forwards and backwards.
// Case-insensitive (e.g., "Racecar" is a palindrome).
// Ignore spaces (e.g., "taco cat" is a palindrome).
// Must handle Unicode characters correctly.
func IsPalindrome(s string) bool {
	return false // Fix me
}

// TODO 3: Implement BuildGreeting
// Given a slice of names, it should return a single string:
// "Hello Name1, Name2, and Name3!"
// If the slice is empty, return "Hello there!"
// If the slice has one name, return "Hello Name1!"
// If the slice has two names, return "Hello Name1 and Name2!"
// MUST use strings.Builder for efficiency.
func BuildGreeting(names []string) string {
	return "" // Fix me
}

// TODO 4: Implement CountVowels
// It should count the number of vowels (a, e, i, o, u) in the string.
// Case-insensitive.
func CountVowels(s string) int {
	return 0 // Fix me
}

// TODO 5: Implement FirstNRunes
// It should return the first N *characters* (runes) of a string.
// If n is greater than the number of runes, return the whole string.
// Do not truncate in the middle of a multi-byte character.
func FirstNRunes(s string, n int) string {
	return "" // Fix me
}
