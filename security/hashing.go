package security

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	DefaultCost = 10
	MinCost     = 4
	MaxCost     = 31
	SaltLength  = 16
	HashLength  = 32
	HashVersion = "v1"
)

var (
	ErrPasswordTooLong           = errors.New("password length exceeds the maximum allowed")
	ErrMismatchedHashAndPassword = errors.New("hashed password does not match the provided password")
	ErrInvalidCost               = errors.New("cost is out of range")
	ErrHashTooShort              = errors.New("hash is too short")
)

type HashedPassword struct {
	Hash       string
	Salt       string
	Cost       int
	HashPrefix string
}

func Hash(password []byte) (*HashedPassword, error) {
	return HashWithCost(password, DefaultCost)
}

func HashWithCost(password []byte, cost int) (*HashedPassword, error) {
	if len(password) > 72 {
		return nil, ErrPasswordTooLong
	}

	if cost < MinCost || cost > MaxCost {
		return nil, ErrInvalidCost
	}

	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}

	hash := deriveKey(password, salt, cost)

	hashString := fmt.Sprintf("%s$%d$%s$%s", HashVersion, cost, salt, hash)

	return &HashedPassword{
		Hash:       hashString,
		Salt:       salt,
		Cost:       cost,
		HashPrefix: HashVersion,
	}, nil
}

func VerifyHash(hashedPassword, password []byte) error {
	hashStr := string(hashedPassword)
	parsedHash, err := parseHash(hashStr)
	if err != nil {
		return err
	}

	derivedHash := deriveKey(password, parsedHash.Salt, parsedHash.Cost)

	if subtle.ConstantTimeCompare([]byte(derivedHash), []byte(parsedHash.Hash)) == 1 {
		return nil
	}

	return ErrMismatchedHashAndPassword
}

func generateSalt() (string, error) {
	salt := make([]byte, SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(salt), nil
}

func deriveKey(password []byte, salt string, cost int) string {
	iterations := 1 << uint(cost)
	hash := hmac.New(sha256.New, []byte(password))
	hash.Write([]byte(salt))
	hashSum := hash.Sum(nil)

	for i := 0; i < iterations-1; i++ {
		hash = hmac.New(sha256.New, hashSum)
		hashSum = hash.Sum(nil)
	}

	return fmt.Sprintf("%x", hashSum)
}

func parseHash(hash string) (*HashedPassword, error) {
	parts := strings.Split(hash, "$")
	if len(parts) != 4 {
		return nil, ErrHashTooShort
	}

	version := parts[0]
	cost, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, ErrInvalidCost
	}
	salt := parts[2]
	hashStr := parts[3]

	if version != HashVersion {
		return nil, fmt.Errorf("unsupported hash version: %s", version)
	}
	if cost < MinCost || cost > MaxCost {
		return nil, ErrInvalidCost
	}

	return &HashedPassword{
		Hash:       hashStr,
		Salt:       salt,
		Cost:       cost,
		HashPrefix: version,
	}, nil
}
