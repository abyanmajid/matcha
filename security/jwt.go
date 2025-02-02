package security

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"
)

type Algorithm string

const (
	HS256 Algorithm = "HS256"
	HS384 Algorithm = "HS384"
	HS512 Algorithm = "HS512"
	RS256 Algorithm = "RS256"
	RS384 Algorithm = "RS384"
	RS512 Algorithm = "RS512"
	ES256 Algorithm = "ES256"
	ES384 Algorithm = "ES384"
	ES512 Algorithm = "ES512"
	EdDSA Algorithm = "EdDSA"
)

type JwtClaims map[string]interface{}

type JWT struct {
	Header    map[string]interface{}
	JwtClaims JwtClaims
	Signature []byte
}

func NewJWT(claims JwtClaims) *JWT {
	return &JWT{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": HS256,
		},
		JwtClaims: claims,
	}
}

func (t *JWT) SetAlgorithm(alg Algorithm) {
	t.Header["alg"] = alg
}

func (t *JWT) Sign(key interface{}) (string, error) {
	headerJSON, err := json.Marshal(t.Header)
	if err != nil {
		return "", fmt.Errorf("failed to encode header: %w", err)
	}
	claimsJSON, err := json.Marshal(t.JwtClaims)
	if err != nil {
		return "", fmt.Errorf("failed to encode claims: %w", err)
	}

	header := base64.RawURLEncoding.EncodeToString(headerJSON)
	payload := base64.RawURLEncoding.EncodeToString(claimsJSON)

	signingInput := fmt.Sprintf("%s.%s", header, payload)

	var signature []byte
	switch alg := t.Header["alg"].(Algorithm); alg {
	case HS256, HS384, HS512:
		signature, err = signHMAC(alg, signingInput, key.([]byte))
	case RS256, RS384, RS512:
		signature, err = signRSA(alg, signingInput, key.(*rsa.PrivateKey))
	case ES256, ES384, ES512:
		signature, err = signECDSA(alg, signingInput, key.(*ecdsa.PrivateKey))
	case EdDSA:
		signature, err = signEd25519(signingInput, key.(ed25519.PrivateKey))
	default:
		return "", fmt.Errorf("unsupported algorithm: %s", alg)
	}
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	sig := base64.RawURLEncoding.EncodeToString(signature)

	return fmt.Sprintf("%s.%s.%s", header, payload, sig), nil
}

func VerifyJWT(token string, key interface{}) (*JWT, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("failed to decode header: %w", err)
	}
	var header map[string]interface{}
	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return nil, fmt.Errorf("failed to parse header: %w", err)
	}

	claimsJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode claims: %w", err)
	}
	var claims JwtClaims
	if err := json.Unmarshal(claimsJSON, &claims); err != nil {
		return nil, fmt.Errorf("failed to parse claims: %w", err)
	}

	signingInput := fmt.Sprintf("%s.%s", parts[0], parts[1])
	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, fmt.Errorf("failed to decode signature: %w", err)
	}

	switch alg := header["alg"].(string); alg {
	case "HS256", "HS384", "HS512":
		err = verifyHMAC(Algorithm(alg), signingInput, signature, key.([]byte))
	case "RS256", "RS384", "RS512":
		err = verifyRSA(Algorithm(alg), signingInput, signature, key.(*rsa.PublicKey))
	case "ES256", "ES384", "ES512":
		err = verifyECDSA(Algorithm(alg), signingInput, signature, key.(*ecdsa.PublicKey))
	case "EdDSA":
		err = verifyEd25519(signingInput, signature, key.(ed25519.PublicKey))
	default:
		return nil, fmt.Errorf("unsupported algorithm: %s", alg)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to verify signature: %w", err)
	}

	if err := validateClaims(claims); err != nil {
		return nil, fmt.Errorf("invalid claims: %w", err)
	}

	return &JWT{
		Header:    header,
		JwtClaims: claims,
		Signature: signature,
	}, nil
}

func validateClaims(claims JwtClaims) error {
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return errors.New("token has expired")
		}
	}

	if iat, ok := claims["iat"].(float64); ok {
		if time.Now().Unix() < int64(iat) {
			return errors.New("token issued in the future")
		}
	}

	return nil
}

func signHMAC(alg Algorithm, input string, key []byte) ([]byte, error) {
	var hash crypto.Hash
	switch alg {
	case HS256:
		hash = crypto.SHA256
	case HS384:
		hash = crypto.SHA384
	case HS512:
		hash = crypto.SHA512
	default:
		return nil, fmt.Errorf("unsupported HMAC algorithm: %s", alg)
	}

	mac := hmac.New(hash.New, key)
	mac.Write([]byte(input))
	return mac.Sum(nil), nil
}

func verifyHMAC(alg Algorithm, input string, signature, key []byte) error {
	expectedSig, err := signHMAC(alg, input, key)
	if err != nil {
		return err
	}
	if !hmac.Equal(signature, expectedSig) {
		return errors.New("invalid signature")
	}
	return nil
}

func signRSA(alg Algorithm, input string, key *rsa.PrivateKey) ([]byte, error) {
	var hash crypto.Hash
	switch alg {
	case RS256:
		hash = crypto.SHA256
	case RS384:
		hash = crypto.SHA384
	case RS512:
		hash = crypto.SHA512
	default:
		return nil, fmt.Errorf("unsupported RSA algorithm: %s", alg)
	}

	hasher := hash.New()
	hasher.Write([]byte(input))
	hashed := hasher.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, key, hash, hashed)
}

func verifyRSA(alg Algorithm, input string, signature []byte, key *rsa.PublicKey) error {
	var hash crypto.Hash
	switch alg {
	case RS256:
		hash = crypto.SHA256
	case RS384:
		hash = crypto.SHA384
	case RS512:
		hash = crypto.SHA512
	default:
		return fmt.Errorf("unsupported RSA algorithm: %s", alg)
	}

	hasher := hash.New()
	hasher.Write([]byte(input))
	hashed := hasher.Sum(nil)

	return rsa.VerifyPKCS1v15(key, hash, hashed, signature)
}

func signECDSA(alg Algorithm, input string, key *ecdsa.PrivateKey) ([]byte, error) {
	var hash crypto.Hash
	switch alg {
	case ES256:
		hash = crypto.SHA256
	case ES384:
		hash = crypto.SHA384
	case ES512:
		hash = crypto.SHA512
	default:
		return nil, fmt.Errorf("unsupported ECDSA algorithm: %s", alg)
	}

	hasher := hash.New()
	hasher.Write([]byte(input))
	hashed := hasher.Sum(nil)

	r, s, err := ecdsa.Sign(rand.Reader, key, hashed)
	if err != nil {
		return nil, err
	}

	sig := append(r.Bytes(), s.Bytes()...)
	return sig, nil
}

func verifyECDSA(alg Algorithm, input string, signature []byte, key *ecdsa.PublicKey) error {
	var hash crypto.Hash
	switch alg {
	case ES256:
		hash = crypto.SHA256
	case ES384:
		hash = crypto.SHA384
	case ES512:
		hash = crypto.SHA512
	default:
		return fmt.Errorf("unsupported ECDSA algorithm: %s", alg)
	}

	hasher := hash.New()
	hasher.Write([]byte(input))
	hashed := hasher.Sum(nil)

	r := new(big.Int).SetBytes(signature[:len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])

	if !ecdsa.Verify(key, hashed, r, s) {
		return errors.New("invalid signature")
	}
	return nil
}

func signEd25519(input string, key ed25519.PrivateKey) ([]byte, error) {
	return ed25519.Sign(key, []byte(input)), nil
}

func verifyEd25519(input string, signature []byte, key ed25519.PublicKey) error {
	if !ed25519.Verify(key, []byte(input), signature) {
		return errors.New("invalid signature")
	}
	return nil
}
