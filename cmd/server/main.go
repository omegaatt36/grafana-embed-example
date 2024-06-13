package main

import (
	_ "embed"
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

func getIndexHtml() (*template.Template, error) {
	index, err := template.New("index").Parse(indexHTML)
	if err != nil {
		return nil, fmt.Errorf("failed to parse index.html: %w", err)
	}

	return index, nil
}

func generateJWT(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user": jwt.MapClaims{
			"email": "viewer@kryptogo.com",
			"name":  "viewer",
		},
		"sub":  "viewer@kryptogo.com",
		"role": "Viewer",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})

	token.Header["kid"] = *keyID

	tokenString, err := token.SignedString([]byte(*secretKey))
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

func auditLog(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL)
		handler(w, r)
	})
}

var grafanaDashboardURL *string = flag.String("grafana-dashboard-url", "", "grafana dashboard url")
var secretKey *string = flag.String("secret-key", "", "secret key")
var keyID *string = flag.String("key-id", "", "key-id")

func main() {
	flag.Parse()

	if secretKey == nil || *secretKey == "" {
		log.Fatal("secret-key is required")
	}
	if grafanaDashboardURL == nil || *grafanaDashboardURL == "" {
		log.Fatal("grafana-dashboard-url is required")
	}
	if keyID == nil || *keyID == "" {
		log.Fatal("key-id is required")
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /token", auditLog(generateJWT))
	router.HandleFunc("GET /", auditLog(serveIndex))

	fmt.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println(err)
	}

	fmt.Println("exiting...")
}
