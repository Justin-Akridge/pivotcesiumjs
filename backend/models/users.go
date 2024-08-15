package models

import (
  "fmt"
  "github.com/pivot/database"
  _ "github.com/lib/pq"
  "github.com/google/uuid"
)

type LoginRequest struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

type User struct {
	ID          uuid.UUID `json:"id"`
  CompanyId   uuid.UUID `json:"company_id"`
  Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
  Permissions string    `json:"permissions"`
}

func StoreUserInDb(email, hashedPassword string) (error) {
  query := `INSERT INTO users (email, password) VALUES ($1, $2)`

  stmt, err := database.DB.Prepare(query)
  if err != nil {
    return fmt.Errorf("failed to prepare statement: %v", err)
  }

  defer stmt.Close()

  _, err = stmt.Exec(email, hashedPassword)
  if err != nil {
    return fmt.Errorf("failed to execute statement", err)
  }

  return nil
}

func GetUserByEmail(email string) (User, error) {
  var user User
  query := "SELECT id, company_id, name, email, password, role, permissions FROM users WHERE email = $1"
  err := database.DB.QueryRow(query, email).Scan(&user.ID, &user.CompanyId, &user.Name, &user.Email, &user.Password, &user.Role, &user.Permissions)
  if err != nil {
    return User{}, err
  }
  return user, nil
}


