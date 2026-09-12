package main

import (
	"testing"
	"golang.org/x/crypto/bcrypt"
)

func TestHashSHA256(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"world", "486ea46224d1bb4fb680f34f7c9ad96a8f24ec88be73ea8e5a6c65260e9cb8a7"},
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := HashSHA256(tt.input); got != tt.want {
				t.Errorf("HashSHA256(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestHashAndCheckPassword(t *testing.T) {
	password := "my_secure_password"
	
	// 1. Test Hashing
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}
	if hash == "" {
		t.Fatalf("HashPassword returned empty string")
	}
	if hash == password {
		t.Fatalf("HashPassword returned plaintext password")
	}

	// Double check that it's a valid bcrypt hash
	_, err = bcrypt.Cost([]byte(hash))
	if err != nil {
		t.Fatalf("Hash is not a valid bcrypt hash: %v", err)
	}

	// 2. Test Checking (Success)
	if !CheckPassword(hash, password) {
		t.Error("CheckPassword failed for correct password")
	}

	// 3. Test Checking (Failure)
	if CheckPassword(hash, "wrong_password") {
		t.Error("CheckPassword succeeded for incorrect password")
	}
}
