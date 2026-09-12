package main

// TODO 1: Implement HashSHA256
// It should take a string and return its SHA-256 hash as a hex-encoded string.
// Hint: use hex.EncodeToString
func HashSHA256(data string) string {
	return "" // Fix me
}

// TODO 2: Implement HashPassword
// It should take a plaintext password and return a bcrypt hash (as a string)
// using bcrypt.DefaultCost.
func HashPassword(password string) (string, error) {
	return "", nil // Fix me
}

// TODO 3: Implement CheckPassword
// It should take a bcrypt hash and a plaintext password.
// Return true if they match, false otherwise.
func CheckPassword(hash string, password string) bool {
	return false // Fix me
}
