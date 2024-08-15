package middleware

import (
  "fmt"
  "strings"
  "context"
  "net/http"
  "github.com/pivot/utils"
)

type contextKey string

const ClaimsContextKey = contextKey("claims")

func AuthMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    publicRoutes := []string{"/login", "/stylesheets/", "/scripts/"}

		// Check if the request path starts with any of the public routes
		for _, route := range publicRoutes {
			if strings.HasPrefix(r.URL.Path, route) {
				next.ServeHTTP(w, r)
				return
			}
		}
      // Extract token from Authorization header
    tokenStr := r.Header.Get("Authorization")
    if tokenStr == "" {
      fmt.Println("Missing Authorization header")
      http.Redirect(w, r, "/login", http.StatusFound)
      http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
      return
    }

    // Remove "Bearer " prefix if present
    if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
      tokenStr = tokenStr[7:]
    }

    // Validate token
    claims, err := utils.ValidateToken(tokenStr)
    fmt.Println(claims)
    if err != nil {
      fmt.Println("Token validation failed:", err)
      http.Redirect(w, r, "/login", http.StatusFound)
      http.Error(w, err.Error(), http.StatusUnauthorized)
      return
    }

    // Add claims to context
    ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}


