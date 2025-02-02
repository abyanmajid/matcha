package main

import (
	"fmt"

	"github.com/abyanmajid/matcha/security"
)

func main() {
	secretKey := []byte("12345678901234567890123456789012")
	plaintext := []byte("Hello, world!")

	ciphertext, err := security.Encrypt(plaintext, secretKey, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted Data:", security.EncodeBase64(ciphertext))

	decryptedText, err := security.Decrypt(ciphertext, secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted Data:", string(decryptedText))
}
