package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

// EncryptSymmetric encrypts data using AES-256-GCM with a user-provided key and IV.
// If no IV is provided, a random one is generated.
func EncryptSymmetric(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher block: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	if len(iv) == 0 {
		iv = make([]byte, gcm.NonceSize())
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return nil, fmt.Errorf("failed to generate IV: %w", err)
		}
	} else if len(iv) != gcm.NonceSize() {
		return nil, fmt.Errorf("invalid IV length: expected %d, got %d", gcm.NonceSize(), len(iv))
	}

	ciphertext := gcm.Seal(nil, iv, plaintext, nil)

	result := append(iv, ciphertext...)
	return result, nil
}

// DecryptSymmetric decrypts data using AES-256-GCM with a user-provided key.
// The IV is extracted from the beginning of the ciphertext.
func DecryptSymmetric(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher block: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	iv, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}

	return plaintext, nil
}

// GenerateRSAKeyPair generates a new RSA key pair.
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate RSA key pair: %w", err)
	}
	return privateKey, &privateKey.PublicKey, nil
}

// EncryptRSA encrypts data using RSA-OAEP.
func EncryptRSA(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, plaintext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt with RSA: %w", err)
	}
	return ciphertext, nil
}

// DecryptRSA decrypts data using RSA-OAEP.
func DecryptRSA(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt with RSA: %w", err)
	}
	return plaintext, nil
}

// EncryptHybrid encrypts data using hybrid encryption (AES + RSA).
// The AES key is encrypted with RSA, and the IV is stored alongside the ciphertext.
func EncryptHybrid(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	// Generate a random AES key
	aesKey := make([]byte, 32) // AES-256
	if _, err := io.ReadFull(rand.Reader, aesKey); err != nil {
		return nil, fmt.Errorf("failed to generate AES key: %w", err)
	}

	ciphertext, err := EncryptSymmetric(plaintext, aesKey, nil) // Generate a random IV
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt with AES: %w", err)
	}

	encryptedKey, err := EncryptRSA(aesKey, publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt AES key with RSA: %w", err)
	}

	encryptedData := append(encryptedKey, ciphertext...)
	return encryptedData, nil
}

// DecryptHybrid decrypts data using hybrid encryption (AES + RSA).
func DecryptHybrid(encryptedData []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	keySize := privateKey.Size()
	if len(encryptedData) < keySize {
		return nil, errors.New("invalid encrypted data")
	}

	encryptedKey := encryptedData[:keySize]
	ciphertext := encryptedData[keySize:]

	aesKey, err := DecryptRSA(encryptedKey, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt AES key with RSA: %w", err)
	}

	plaintext, err := DecryptSymmetric(ciphertext, aesKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt with AES: %w", err)
	}

	return plaintext, nil
}

// SerializeRSAPrivateKey serializes an RSA private key to PEM format.
func SerializeRSAPrivateKey(privateKey *rsa.PrivateKey) []byte {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	return privateKeyPEM
}

// DeserializeRSAPrivateKey deserializes an RSA private key from PEM format.
func DeserializeRSAPrivateKey(privateKeyPEM []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid PEM block")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}
	return privateKey, nil
}

// SerializeRSAPublicKey serializes an RSA public key to PEM format.
func SerializeRSAPublicKey(publicKey *rsa.PublicKey) []byte {
	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return publicKeyPEM
}

// DeserializeRSAPublicKey deserializes an RSA public key from PEM format.
func DeserializeRSAPublicKey(publicKeyPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("invalid PEM block")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	return publicKey, nil
}
