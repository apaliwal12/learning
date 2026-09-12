package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// ─────── 3. SYMMETRIC ENCRYPTION (AES-GCM) ───────
// Symmetric encryption uses the SAME key to encrypt and decrypt data.
// AES (Advanced Encryption Standard) is the industry standard.
// ALWAYS use Authenticated Encryption with Associated Data (AEAD) like GCM!
// Raw AES or AES-CBC does not verify that the ciphertext hasn't been tampered with.

func demonstrateEncryption() {
	fmt.Println("\n--- AES-GCM (Symmetric Encryption) ---")
	
	// 1. The Key
	// AES keys must be either 16, 24, or 32 bytes (for AES-128, AES-192, or AES-256).
	// In production, NEVER hardcode this. Load it from a secret manager or env var.
	key := []byte("this_is_a_32_byte_super_secret_k") // 32 bytes
	plaintext := []byte("This is a highly confidential message.")
	
	// 2. Create the Cipher Block
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return
	}
	
	// 3. Wrap it in GCM (Galois/Counter Mode)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Error creating GCM:", err)
		return
	}
	
	// 4. Create a Nonce (Number used ONCE)
	// NEVER reuse a nonce with the same key. If you do, AES-GCM breaks catastrophically.
	// The standard way is to generate it randomly.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("Error generating nonce:", err)
		return
	}
	
	// 5. Encrypt (Seal)
	// Seal appends the encrypted data to the first argument.
	// A common pattern is to prepend the nonce to the ciphertext, so you have it
	// available for decryption later. The nonce does NOT need to be kept secret!
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	fmt.Printf("Encrypted (hex): %x\n", ciphertext)
	
	// ─────── DECRYPTION ───────
	
	// 1. Extract the nonce and the actual ciphertext
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println("Ciphertext too short")
		return
	}
	
	extractedNonce, actualCiphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	
	// 2. Decrypt (Open)
	// If the ciphertext was tampered with, Open will return an error!
	decryptedMsg, err := gcm.Open(nil, extractedNonce, actualCiphertext, nil)
	if err != nil {
		fmt.Println("Decryption failed (tampering detected?):", err)
		return
	}
	
	fmt.Printf("Decrypted: %s\n", decryptedMsg)
}
