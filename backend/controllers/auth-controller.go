package controllers

import (
  "path/filepath"
  "database/sql"
  "encoding/json"
  "log"
  "fmt"
  "net/http"
  "github.com/pivot/models"
  "github.com/pivot/utils"
)

func HomePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
      case http.MethodGet:
        ServeHomePage(w, r)

      //case http.MethodPost:
      //  handleLoginRequest(w, r)

      default:
        fmt.Fprintln(w, "method not allowed %s", r.Method)
    }

  }
}
func ServeHomePage(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, filepath.Join("dist", "index.html"))
}

func HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
      case http.MethodGet:
        ServeLoginPage(w, r)

      case http.MethodPost:
        handleLoginRequest(w, r)

      default:
        fmt.Fprintln(w, "method not allowed %s", r.Method)
    }
	}
}

func ServeLoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../frontend/dist/login.html")
}

func handleLoginRequest(w http.ResponseWriter, r *http.Request) {
  var req models.LoginRequest
  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
      http.Error(w, "failed to parse login information", http.StatusBadRequest)
      return
  }

  if req.Email == "" || req.Password == "" {
      http.Error(w, "Email and password cannot be blank", http.StatusBadRequest)
      return
  }

  user, err := models.GetUserByEmail(req.Email)
  if err != nil {
    log.Printf("Error querying user: %v", err)
    if err == sql.ErrNoRows {
      http.Error(w, "user does not exist", http.StatusUnauthorized)
      return
    } else {
      http.Error(w, "Database error", http.StatusInternalServerError)
      return
    }
  }

  if !utils.CheckPasswordHash(req.Password, user.Password) {
      http.Error(w, "Invalid email or password", http.StatusUnauthorized)
      return
  }

  // Generate a token
  token, err := utils.GenerateToken(user.Email, user.Role, user.Permissions, user.CompanyId.String(), user.Name)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }

  claims, err := utils.ValidateToken(token)
  if err != nil {
      http.Error(w, "Failed to parse token claims", http.StatusInternalServerError)
      return
  }

  http.SetCookie(w, &http.Cookie{
      Name:     "authToken",
      Value:    token,
      HttpOnly: true,
      Secure:   true,
      SameSite: http.SameSiteStrictMode,
      Path:     "/",
  })

  response := map[string]interface{}{
      "token":   token,
      "claims": claims,
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(response)
}

//func handleLoginRequest(w http.ResponseWriter, r *http.Request) {
//  var req models.LoginRequest
//  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//    http.Error(w, "failed to parse login information", http.StatusBadRequest)
//    return
//  }
//
//  if req.Email == "" || req.Password == "" {
//    http.Error(w, "Email and password cannot be blank", http.StatusBadRequest)
//    return
//  }
//
//  var user models.User
//  query := "SELECT id, company_id, name, email, password, role, permissions FROM users WHERE email = $1"
//  err := database.DB.QueryRow(query, req.Email).Scan(&user.ID, &user.CompanyId, &user.Name, &user.Email, &user.Password, &user.Role, &user.Permissions)
//  if err != nil {
//    log.Printf("Error querying user: %v", err)
//    if err == sql.ErrNoRows {
//      http.Error(w, "user does not exist", http.StatusUnauthorized)
//      return
//    } else {
//      http.Error(w, "Database error", http.StatusInternalServerError)
//      return
//    }
//  }
//
//  if !utils.CheckPasswordHash(req.Password, user.Password) {
//    http.Error(w, "Invalid email or password", http.StatusUnauthorized)
//    return
//  } else {
//    fmt.Println("successful login")
//  }
//
//  // Generate a token
//  companyIdStr := user.CompanyId.String()
//  token, err := utils.GenerateToken(user.Email, user.Role, user.Permissions, companyIdStr, user.Name)
//  if err != nil {
//    http.Error(w, err.Error(), http.StatusInternalServerError)
//    return
//  }
//
//  claims, err := utils.ValidateToken(token)
//  if err != nil {
//    http.Error(w, "Failed to parse token claims", http.StatusInternalServerError)
//    return
//  }
//
//  http.SetCookie(w, &http.Cookie{
//    Name:     "authToken",
//    Value:    token,
//    HttpOnly: true,
//    Secure:   true,
//    SameSite: http.SameSiteStrictMode,
//    Path:     "/",
//  })
//
//  response := map[string]interface{}{
//    "token":   token,
//    "claims": claims,
//  }
//
//  w.WriteHeader(http.StatusOK)
//  json.NewEncoder(w).Encode(response)
//}
func VerifyTokenHandler() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/", http.StatusSeeOther)
  }
}

//func VerifyTokenHandler() http.HandlerFunc {
//  return func(w http.ResponseWriter, r *http.Request) {
//    fmt.Println("HERE")
//    // Extract token from Authorization header
//    authHeader := r.Header.Get("Authorization")
//    if authHeader == "" {
//      http.Error(w, "Authorization header is required", http.StatusBadRequest)
//      return
//    }
//
//    // Token should be in the format "Bearer <token>"
//    tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
//    if tokenStr == "" {
//      http.Error(w, "Bearer token is required", http.StatusBadRequest)
//      return
//    }
//
//    // Validate the token
//    claims, err := utils.ValidateToken(tokenStr)
//    if err != nil {
//      http.Error(w, err.Error(), http.StatusUnauthorized)
//      return
//    }
//
//    // Respond with token validity and user claims
//    response := map[string]interface{}{
//      "valid":  true,
//      "claims": claims,
//    }
//
//    w.Header().Set("Content-Type", "application/json")
//    w.WriteHeader(http.StatusOK)
//    json.NewEncoder(w).Encode(response)
//  }
//}

func RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

    // register and login take same fields
    //var req models.LoginRequest
    //if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    //  http.Error(w, "Bad Request", http.StatusBadRequest)
    //  return
    //}

    //hashedPassword, err := utils.HashPassword(req.Password)
    //if err != nil {
    //  http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    //  return
    //}

    ////err = models.StoreUserInDb(req.Email, hashedPassword)
    //if err != nil {
    //  log.Printf("Error storing user: %v", err)
    //  http.Error(w, "failed to store user in database", http.StatusInternalServerError)
    //  return
    //}

    //response := map[string]bool{"authenticated": true}
    //w.Header().Set("Content-Type", "application/json")
    //json.NewEncoder(w).Encode(response)
  }
}

