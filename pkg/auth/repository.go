package auth

// UserRepository interface
type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}
