package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRsaSign(t *testing.T) {
	publicKey, privateKey := generateRsaKeyPair()

	data := []byte("Hello")
	signature := rsaSign(data, privateKey)
	fmt.Println("Signature:", signature)

	valid := rsaSignVerify(data, signature, publicKey)
	fmt.Println("Signature valid:", valid)
}

// sian demo: rsa2048 + sha256
// rsaSign signs the data with a private key
func rsaSign(data []byte, privateKey *rsa.PrivateKey) string {
	hashed := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	// PSS 填充方式：使用随机数填充，更新更安全
	// signature, err = rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], nil)
	return hex.EncodeToString(signature)
}

// rsaSignVerify verifies the signature with a public key
func rsaSignVerify(data []byte, signature string, publicKey *rsa.PublicKey) bool {
	signatureBytes, _ := hex.DecodeString(signature)
	hashed := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signatureBytes)
	if err != nil {
		fmt.Println("Signature verification failed.")
		return false
	}
	return true
}
