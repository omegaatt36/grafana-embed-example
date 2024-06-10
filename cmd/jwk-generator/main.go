package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	err := createKeys()
	if err != nil {
		fmt.Printf("Error creating keys: %v\n", err)
	}
}

func createKeys() error {
	keyID := uuid.New().String()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("failed to generate private key: %v", err)
	}

	privKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return fmt.Errorf("failed to marshal public key: %v", err)
	}
	pubKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubKeyBytes,
		},
	)

	err = os.WriteFile("rsa.pem", privKeyPEM, 0644)
	if err != nil {
		return fmt.Errorf("failed to write private key to file: %v", err)
	}

	err = os.WriteFile("rsa_pub.pem", pubKeyPEM, 0644)
	if err != nil {
		return fmt.Errorf("failed to write public key to file: %v", err)
	}

	fmt.Printf("Keys created successfully with key ID: %s\n", keyID)
	return nil
}
