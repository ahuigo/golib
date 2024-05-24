package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

func TestRsaEncrypt(t *testing.T) {
	publicKey, privateKey := generateRsaKeyPair()
	plainText := "Hello, World!"
	encryptedText, err := encryptWithPublicKey([]byte(plainText), publicKey)
	if err != nil {
		log.Fatalf("Failed to encrypt text: %s", err)
	}

	encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedText)
	fmt.Printf("\nPlain text: %s\n", plainText)
	fmt.Printf("Encrypted text (base64): %s\n", encryptedBase64)

	decryptedText, err := decryptWithPrivateKey(encryptedText, privateKey)
	if err != nil {
		log.Fatalf("Failed to decrypt text: %s", err)
	}

	fmt.Printf("Decrypted text: %s\n", decryptedText)
}

func encryptWithPublicKey(plainText []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	encryptedText, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		plainText,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("encryption failed: %s", err)
	}

	return encryptedText, nil
}

func decryptWithPrivateKey(encryptedText []byte, privateKey *rsa.PrivateKey) (string, error) {
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
