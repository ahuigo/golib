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

func TestRsaEncrypt(t *testing.T) {
	publicKeyPem, privateKeyPem, err := generateRsaKeyPemPair()
	if err != nil {
		log.Fatalf("Failed to generate key: %s", err)
	}

	fmt.Printf("Public Key (PEM):\n%s\n", publicKeyPem)
	// fmt.Printf("Private Key:\n%s", privateKey)

	plainText := "Hello, World!"
	encryptedText, err := encryptWithPublicKey([]byte(plainText), publicKeyPem)
	if err != nil {
		log.Fatalf("Failed to encrypt text: %s", err)
	}

	encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedText)
	fmt.Printf("\nPlain text: %s\n", plainText)
	fmt.Printf("Encrypted text (base64): %s\n", encryptedBase64)

	decryptedText, err := decryptWithPrivateKey(encryptedText, privateKeyPem)
	if err != nil {
		log.Fatalf("Failed to decrypt text: %s", err)
	}

	fmt.Printf("Decrypted text: %s\n", decryptedText)
}

func encryptWithPublicKey(plainText, publicKeyPem []byte) ([]byte, error) {
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

func decryptWithPrivateKey(encryptedText []byte, privateKeyPem []byte) (string, error) {
	// Assume privateKeyPem is your PEM-encoded private key
	block, _ := pem.Decode([]byte(privateKeyPem))
	if block == nil {
		panic("failed to parse PEM block containing the private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic("failed to parse private key: " + err.Error())
	}

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
