package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"strings"
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
	// 1. private key:  PKCS#1, ASN.1 DER 形式的字节
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPem = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// 2. public key
	/**
	1. PKIX 是一个公钥基础设施 (PKI) 的标准，它定义了证书和公钥的格式。PKIX 格式的公钥通常包含在 X.509 证书中(python的: PKCS1_OAEP)
	2. PKCS#1 是 RSA 加密标准的一部分，它定义了 RSA 公钥和私钥的格式。
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	*/
	publicKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	publicKeyPem = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	fmt.Printf("publicKeyPem:\n%s\n", publicKeyPem)
	fmt.Printf("privateKeyPem:\n%s\n", privateKeyPem)
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
	if block == nil || !strings.HasSuffix(block.Type, "PUBLIC KEY") {
		panic("invalid public key")
	}

	rsaPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil && strings.Contains(err.Error(), "use ParsePKIXPublicKey") {
		var ok bool
		var publicKey1 any
		publicKey1, err = x509.ParsePKIXPublicKey(block.Bytes)
		if err == nil {
			if rsaPublicKey, ok = publicKey1.(*rsa.PublicKey); !ok {
				err = fmt.Errorf("not an RSA public key")
			}
		}

	}
	if err != nil {
		panic("not an RSA public key")
	}
	return rsaPublicKey, privateKey
}
