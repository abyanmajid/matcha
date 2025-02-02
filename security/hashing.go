package security

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

type HashAlgorithm string

const (
	SHA256   HashAlgorithm = "sha256"
	SHA3_256 HashAlgorithm = "sha3-256"
	BLAKE2b  HashAlgorithm = "blake2b"
)

type Hasher interface {
	Hash(data []byte) []byte
	HashString(data string) string
}

type hasherImpl struct {
	algorithm HashAlgorithm
}

func NewHasher(algorithm HashAlgorithm) Hasher {
	return &hasherImpl{
		algorithm: algorithm,
	}
}

func (h *hasherImpl) Hash(data []byte) []byte {
	var hashFunc hash.Hash

	switch h.algorithm {
	case SHA256:
		hashFunc = sha256.New()
	case SHA3_256:
		hashFunc = sha3.New256()
	case BLAKE2b:
		hashFunc, _ = blake2b.New256(nil)
	default:
		panic("unsupported hashing algorithm")
	}

	hashFunc.Write(data)
	return hashFunc.Sum(nil)
}

func (h *hasherImpl) HashString(data string) string {
	hashBytes := h.Hash([]byte(data))
	return fmt.Sprintf("%x", hashBytes)
}

type KeyedHasher interface {
	Hash(data []byte, key []byte) []byte
	HashString(data string, key string) string
}

type keyedHasherImpl struct {
	algorithm HashAlgorithm
}

func NewKeyedHasher(algorithm HashAlgorithm) KeyedHasher {
	return &keyedHasherImpl{
		algorithm: algorithm,
	}
}

func (h *keyedHasherImpl) Hash(data []byte, key []byte) []byte {
	var hashFunc func() hash.Hash

	switch h.algorithm {
	case SHA256:
		hashFunc = sha256.New
	case SHA3_256:
		hashFunc = sha3.New256
	case BLAKE2b:
		hashFunc = func() hash.Hash {
			h, _ := blake2b.New256(key)
			return h
		}
	default:
		panic("unsupported hashing algorithm")
	}

	mac := hmac.New(hashFunc, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func (h *keyedHasherImpl) HashString(data string, key string) string {
	hashBytes := h.Hash([]byte(data), []byte(key))
	return fmt.Sprintf("%x", hashBytes)
}

type SaltedHasher interface {
	Hash(data []byte) ([]byte, []byte)
	HashString(data string) (string, string)
}

type saltedHasherImpl struct {
	algorithm HashAlgorithm
}

func NewSaltedHasher(algorithm HashAlgorithm) SaltedHasher {
	return &saltedHasherImpl{
		algorithm: algorithm,
	}
}

func (h *saltedHasherImpl) Hash(data []byte) ([]byte, []byte) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	var hashFunc hash.Hash

	switch h.algorithm {
	case SHA256:
		hashFunc = sha256.New()
	case SHA3_256:
		hashFunc = sha3.New256()
	case BLAKE2b:
		hashFunc, _ = blake2b.New256(nil)
	default:
		panic("unsupported hashing algorithm")
	}

	hashFunc.Write(salt)
	hashFunc.Write(data)
	return hashFunc.Sum(nil), salt
}

func (h *saltedHasherImpl) HashString(data string) (string, string) {
	hashBytes, salt := h.Hash([]byte(data))
	return fmt.Sprintf("%x", hashBytes), base64.StdEncoding.EncodeToString(salt)
}

func Hash(data string) string {
	hasher := NewHasher(SHA256)
	return hasher.HashString(data)
}

func VerifyHash(data string, hash string) bool {
	return Hash(data) == hash
}

func HashWithKey(data string, key string) string {
	keyedHasher := NewKeyedHasher(SHA256)
	return keyedHasher.HashString(data, key)
}

func HashWithSalt(data string) (string, string) {
	saltedHasher := NewSaltedHasher(SHA256)
	return saltedHasher.HashString(data)
}

func VerifyHashWithSalt(data string, hash string, salt string) bool {
	saltedHasher := NewSaltedHasher(SHA256)
	computedHash, _ := saltedHasher.HashString(data)
	return computedHash == hash
}
