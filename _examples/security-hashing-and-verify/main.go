package main

import (
	"fmt"
	"log"

	"github.com/abyanmajid/matcha/security"
)

func main() {
	// Usecase 1: Hash a password with default cost
	password := []byte("mySuperSecretPassword")
	hashedPassword, err := security.Hash(password)
	if err != nil {
		log.Fatal("Error hashing password:", err)
	}

	fmt.Println("Hashed password with default cost:", hashedPassword.Hash)

	// Usecase 2: Hash a password with a custom cost
	customCost := 12
	hashedPasswordWithCost, err := security.HashWithCost(password, customCost)
	if err != nil {
		log.Fatal("Error hashing password with custom cost:", err)
	}

	fmt.Println("Hashed password with custom cost:", hashedPasswordWithCost.Hash)

	// Usecase 3: Verify the hashed password
	err = security.VerifyHash([]byte(hashedPassword.Hash), password)
	if err != nil {
		log.Fatal("Password verification failed:", err)
	} else {
		fmt.Println("Password verification succeeded.")
	}

	incorrectPassword := []byte("wrongPassword")
	err = security.VerifyHash([]byte(hashedPassword.Hash), incorrectPassword)
	if err != nil {
		fmt.Println("Password verification failed as expected:", err)
	} else {
		fmt.Println("Password verification unexpectedly succeeded.")
	}
}
