package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// While HMAC uses a secret key, this is not the same as encryption
	// and cannot be used to decrypt a password.
	secretKeyForHash := "secret-key"
	password := "this is a totally secret password nobody will guess"
	// Setup our hashing function
	h := hmac.New(sha256.New, []byte(secretKeyForHash))
	// Write data to our hashing function
	h.Write([]byte(password))
	// Get the resulting hash
	result := h.Sum(nil)
	// The resulting hash is binary, so encode it to hex so we can read it as letters and numbers
	fmt.Println(hex.EncodeToString(result))
}
