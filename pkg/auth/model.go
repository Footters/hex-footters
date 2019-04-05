package auth

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}
