package variables

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

const (
	// SigningAlgorithm is the alg used to sign tokens
	SigningAlgorithm = "RS256"
	// EnvTokenSignKey is the location of the key to sign with
	EnvTokenSignKey = "TOKEN_SIGN_KEY"
	// EnvTokenVerifyKey is the location of the key to sign with
	EnvTokenVerifyKey = "TOKEN_VERIFY_KEY"
	// TokenSignKeyDefault is the location of the key to sign with
	TokenSignKeyDefault = "keys/token/private.pem"
	// TokenVerifyKeyDefault is the location of the key to sign with
	TokenVerifyKeyDefault = "keys/token/public.pem"
)

var (
	// TokenSignKey is the key used to sign tokens
	TokenSignKey *rsa.PrivateKey
	// TokenVerifyKey is the key used to verify tokens
	TokenVerifyKey *rsa.PublicKey
	// BackendToken is the token that backend services send each other
	BackendToken string
)

// LoadTokenKeys loads the tokens from files
func LoadTokenKeys() error {
	err := LoadTokenSignKey()
	if err != nil {
		log.Println("Error loading TokenSignKey:", err)
		return err
	}
	err = LoadTokenVerifyKey()
	if err != nil {
		log.Println("Error loading TokenVerifyKey:", err)
		return err
	}
	err = CreateBackendToken()
	if err != nil {
		log.Println("Failed to sign BackendToken:", err)
		return err
	}
	return nil
}

// LoadTokenSignKey reads the private key from EnvTokenSignKey
func LoadTokenSignKey() error {
	tokenSignKeyPath := os.Getenv(EnvTokenSignKey)
	if tokenSignKeyPath == "" {
		tokenSignKeyPath = TokenSignKeyDefault
	}
	tokenSignKeyPem, err := ioutil.ReadFile(tokenSignKeyPath)
	if err != nil {
		return err
	}
	pemBlock, _ := pem.Decode(tokenSignKeyPem)
	if pemBlock.Type != "RSA PRIVATE KEY" {
		return errors.New("Key type was not RSA PRIVATE KEY")
	}
	tokenSignKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return err
	}
	TokenSignKey = tokenSignKey
	return nil
}

// LoadTokenVerifyKey reads the public key from EnvTokenVerifyKey
func LoadTokenVerifyKey() error {
	tokenVerifyKeyPath := os.Getenv(EnvTokenVerifyKey)
	if tokenVerifyKeyPath == "" {
		tokenVerifyKeyPath = TokenVerifyKeyDefault
	}
	tokenVerifyKeyPem, err := ioutil.ReadFile(tokenVerifyKeyPath)
	if err != nil {
		return err
	}
	pemBlock, _ := pem.Decode(tokenVerifyKeyPem)
	if pemBlock.Type != "PUBLIC KEY" {
		return errors.New("Key type was not PUBLIC KEY")
	}
	tokenVerifyKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return err
	}
	var ok bool
	TokenVerifyKey, ok = tokenVerifyKey.(*rsa.PublicKey)
	if !ok {
		return errors.New(tokenVerifyKeyPath + " could not be loaded")
	}
	return nil
}

// CreateBackendToken uses the signing key to create a token
func CreateBackendToken() error {
	var err error
	token := jwt.New(jwt.GetSigningMethod(SigningAlgorithm))

	token.Claims["id"] = "backend"
	token.Claims["backend"] = true

	if TokenSignKey == nil {
		return errors.New("TokenSignKey is nil")
	}
	BackendToken, err = token.SignedString(TokenSignKey)
	if err != nil {
		return err
	}
	return nil
}
