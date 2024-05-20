package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"testing"
)

func TestGenerateRsaKeyPair(t *testing.T) {
	generateRsaKeyPair()
}

func generateRsaKeyPemPair() (privateKeyPem, publicKeyPem []byte, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate key: %s", err)
	}
	// private key
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPem = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	// public key
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatalf("Failed to extract public key: %s", err)
	}
	publicKeyPem = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	fmt.Printf("publicKeyPem:\n%s\n", publicKeyPem)
	return publicKeyPem, privateKeyPem, err
}

func generateRsaKeyPair() (*rsa.PublicKey, *rsa.PrivateKey) {
	publicKeyPem, privateKeyPem, err := generateRsaKeyPemPair()
	if err != nil {
		panic(err)
	}

	// 1. private key
	block, _ := pem.Decode([]byte(privateKeyPem))
	if block == nil {
		panic("failed to parse PEM block containing the private key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic("failed to parse private key: " + err.Error())
	}

	// 2. public key
	block, _ = pem.Decode(publicKeyPem)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		panic("invalid public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse public key: " + err.Error())
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		panic("not an RSA public key")
	}
	return rsaPublicKey, privateKey
}
