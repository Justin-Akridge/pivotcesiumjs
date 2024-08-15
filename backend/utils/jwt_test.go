package utils

import (
  "fmt"
	"os"
	"testing"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken(t *testing.T) {
	// Set a test environment variable for the JWT secret key
	os.Setenv("JWT_SECRET_KEY", "test_secret_key")
	defer os.Unsetenv("JWT_SECRET_KEY") // Clean up environment variable after test

	// Define a test email
	testEmail := "test@example.com"

	// Generate the token
	token, err := GenerateToken(testEmail)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Parse and validate the token
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

  claims, ok := parsedToken.Claims.(*Claims)
	if !ok {
		t.Fatalf("Claims are not of type *Claims")
	}

  fmt.Println(claims)

	// Check the email in the claims
	if claims.Email != testEmail {
		t.Errorf("Expected email %s, got %s", testEmail, claims.Email)
	}

	// Check if the token is expired
	if time.Now().After(claims.ExpiresAt.Time) {
		t.Errorf("Token is expired")
	}

	// Check if the token is issued at a valid time
	if claims.IssuedAt.Time.After(time.Now()) {
		t.Errorf("Token issued at a future time")
	}

  fmt.Println(claims)
}


