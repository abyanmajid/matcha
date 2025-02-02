package main

import (
	"encoding/base64"
	"fmt"

	"github.com/abyanmajid/matcha/security"
)

func main() {
	secretKey := []byte("12345678901234567890123456789012")

	plaintext := []byte("Hello, world!")
	ciphertext, err := security.EncryptSymmetric(plaintext, secretKey, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted Data:", base64.StdEncoding.EncodeToString(ciphertext))

	decryptedText, err := security.DecryptSymmetric(ciphertext, secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted Data:", string(decryptedText))
}
