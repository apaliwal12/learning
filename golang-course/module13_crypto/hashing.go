package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// ─────── 1. FAST HASHING (e.g., Checksums) ───────
// Algorithms like SHA-256 are designed to be FAST.
// They are great for verifying file integrity, digital signatures,
// and deduplicating data.
// THEY ARE TERRIBLE FOR PASSWORDS because an attacker can guess billions
// of passwords per second using GPUs.

func demonstrateSHA256() {
	fmt.Println("\n--- SHA-256 (Fast Hashing) ---")
	
	data := []byte("hello world")
	
	// Create a new hash.Hash (which implements io.Writer)
	h := sha256.New()
	
	// Write data to it
	h.Write(data)
	
	// Get the finalized hash as a byte slice
	// Sum() takes an optional byte slice to append the hash to. Usually we pass nil.
	hashBytes := h.Sum(nil)
	
	// Convert bytes to a hex string for printing
	hashStr := hex.EncodeToString(hashBytes)
	fmt.Printf("SHA-256 of %q: %s\n", data, hashStr)
}

// ─────── 2. SLOW HASHING (Passwords) ───────
// Algorithms like bcrypt, Argon2, and scrypt are designed to be SLOW.
// They purposefully eat up CPU/Memory to prevent brute-force attacks.
// They also automatically generate and incorporate a unique "salt" into every hash.

func demonstrateBcrypt() {
	fmt.Println("\n--- bcrypt (Password Hashing) ---")
	
	password := "super_secret_p4ssw0rd!"
	
	// 1. Hash the password
	// The "Cost" determines how slow the algorithm is. Higher is slower/more secure.
	// bcrypt.DefaultCost is currently 10.
	start := time.Now()
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	fmt.Printf("Hashing took %v\n", time.Since(start))
	fmt.Printf("Hashed Password: %s\n", string(hashedBytes))
	
	// 2. Verify the password
	// Notice we pass the HASH and the RAW PASSWORD.
	// bcrypt extracts the salt from the hash, hashes the raw password with it,
	// and checks if they match.
	err = bcrypt.CompareHashAndPassword(hashedBytes, []byte(password))
	if err == nil {
		fmt.Println("Password match: SUCCESS")
	} else {
		fmt.Println("Password match: FAILED")
	}
	
	// Check a wrong password
	err = bcrypt.CompareHashAndPassword(hashedBytes, []byte("wrong_password"))
	fmt.Println("Wrong password check error:", err)
}
