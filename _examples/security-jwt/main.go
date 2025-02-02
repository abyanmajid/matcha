package main

import (
	"fmt"
	"time"

	"github.com/abyanmajid/matcha/security"
)

func main() {
	token := security.NewJWT(security.JwtClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	key := []byte("your-256-bit-secret")
	signedToken, err := token.Sign(key)
	if err != nil {
		panic(err)
	}
	fmt.Println("Signed Token:", signedToken)

	verifiedToken, err := security.VerifyJWT(signedToken, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("Verified Claims:", verifiedToken.JwtClaims)
}
