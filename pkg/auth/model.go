package auth

// User model
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
