package utils

import (
  "fmt"
  "os"
  "time"
  "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
  Email       string `json:"email"`
  Role        string `json:"role"`
  Permissions string `json:"permissions"`
  Name        string `json:"name"`
  CompanyId   string `json:"companyId"`
  jwt.RegisteredClaims
}

func GenerateToken(email, role, permissions, companyId, name string) (string, error) {
  expirationTime := time.Now().Add(24 * time.Hour)
  claims := &Claims{
    Email: email,
    Role: role,
    CompanyId: companyId,
    Permissions: permissions,
    Name: name,
    RegisteredClaims: jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(expirationTime),
      IssuedAt:  jwt.NewNumericDate(time.Now()),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
  var claims Claims

  // Parse the token
  token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
    // Ensure that the token method matches what you expect
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method")
    }
    return jwtKey, nil
  })
  if err != nil {
    return nil, fmt.Errorf("failed to parse token: %w", err)
  }

  // Check if the token is valid
  if !token.Valid {
    return nil, fmt.Errorf("invalid token")
  }

  // Check if the token is expired
  if claims.ExpiresAt.Time.Before(time.Now()) {
    return nil, fmt.Errorf("token expired")
  }

  return &claims, nil
}

