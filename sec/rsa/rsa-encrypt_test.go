package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"testing"
)

func TestRsaDecrypt(t *testing.T) {
	// Generate an RSA key pair (private + public keys)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate key: %s", err)
	}

	// Extract the public key from the private key
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatalf("Failed to extract public key: %s", err)
	}

	// Encode the public key into PEM format
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	fmt.Printf("Public Key (PEM):\n%s\n", publicKeyPem)

	plainText := "Hello, World!"
	encryptedText, err := EncryptWithPublicKey([]byte(plainText), publicKeyPem)
	if err != nil {
		log.Fatalf("Failed to encrypt text: %s", err)
	}

	encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedText)
	fmt.Printf("\nPlain text: %s\n", plainText)
	fmt.Printf("Encrypted text (base64): %s\n", encryptedBase64)

	decryptedText, err := DecryptWithPrivateKey(encryptedText, privateKey)
	if err != nil {
		log.Fatalf("Failed to decrypt text: %s", err)
	}

	fmt.Printf("Decrypted text: %s\n", decryptedText)
}

func EncryptWithPublicKey(plainText, publicKeyPem []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKeyPem)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("invalid public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %s", err)
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	encryptedText, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPublicKey,
		plainText,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("encryption failed: %s", err)
	}

	return encryptedText, nil
}

func DecryptWithPrivateKey(encryptedText []byte, privateKey *rsa.PrivateKey) (string, error) {
	decryptedBytes, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		encryptedText,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %s", err)
	}

	return string(decryptedBytes), nil
}
