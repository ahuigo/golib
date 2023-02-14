// refer1: https://stackoverflow.com/questions/67625752/how-to-use-a-certificate-from-a-certificate-store-and-run-tls-in-gin-framework-i
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization:  []string{"ORGANIZATION_NAME"},
			Country:       []string{"COUNTRY_CODE"},
			Province:      []string{"PROVINCE"},
			Locality:      []string{"CITY"},
			StreetAddress: []string{"ADDRESS"},
			PostalCode:    []string{"POSTAL_CODE"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	// Sign the certificate
	certificate, _ := x509.CreateCertificate(rand.Reader, cert, cert, pub, priv)

	certBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certificate})
	keyBytes := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	// Generate a key pair from your pem-encoded cert and key ([]byte).
	x509Cert, _ := tls.X509KeyPair(certBytes, keyBytes)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{x509Cert}}
	server := http.Server{Addr: ":3000", Handler: http.HandlerFunc(handler), TLSConfig: tlsConfig}

	err := server.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}
