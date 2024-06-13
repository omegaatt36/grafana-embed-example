package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/square/go-jose.v2"
)

var secretKey *string = flag.String("secret-key", "", "secret key")
var keyID *string = flag.String("key-id", "", "key-id")

func main() {
	flag.Parse()
	if secretKey == nil || *secretKey == "" {
		log.Fatal("secret-key is required")
	}

	if keyID == nil || *keyID == "" {
		log.Fatal("key-id is required")
	}

	rawKey := []byte(*secretKey)

	symKey := jose.JSONWebKey{
		Key:       rawKey,
		KeyID:     *keyID,
		Algorithm: string(jose.HS512),
		Use:       "sig",
	}

	jwkJSON, err := json.MarshalIndent(jose.JSONWebKeySet{Keys: []jose.JSONWebKey{symKey}}, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal JWK: %s\n", err)
		return
	}

	file, err := os.Create("jwks.json")
	if err != nil {
		fmt.Printf("Failed to create file: %s\n", err)
		return
	}
	defer file.Close()

	if _, err := file.Write(jwkJSON); err != nil {
		fmt.Printf("Failed to write JWK to file: %s\n", err)
		return
	}

	fmt.Println("JWK successfully written to jwks.json")
}
