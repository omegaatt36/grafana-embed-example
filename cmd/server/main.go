package main

import (
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//go:embed index.html
var indexHTML string

//go:embed rsa.pem
var privateKey string

func getIndexHtml() (*template.Template, error) {
	index, err := template.New("index").Parse(indexHTML)
	if err != nil {
		return nil, fmt.Errorf("failed to parse index.html: %w", err)
	}

	return index, nil
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		panic("Failed to decode PEM block containing private key")
	}

	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return pk, nil
}

func generateJWT(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user":  "demo",
		"email": "user@demo.com",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})

	pk, err := getPrivateKey()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokenString, err := token.SignedString(pk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte(tokenString)); err != nil {
		log.Println(err)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	index, err := getIndexHtml()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := index.Execute(w, map[string]any{
		"DashboardURL": *grafanaDashboardURL,
	}); err != nil {
		log.Println(err)
	}
}

var grafanaDashboardURL *string = flag.String("grafana-dashboard-url", "", "grafana dashboard url")

func main() {
	flag.Parse()
	if grafanaDashboardURL == nil || *grafanaDashboardURL == "" {
		log.Fatal("grafana-dashboard-url is required")
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /token", generateJWT)
	router.HandleFunc("GET /", serveIndex)

	fmt.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println(err)
	}

	fmt.Println("exiting...")
}
