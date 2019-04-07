package auth

import "errors"

// Service interface
type Service interface {
	RegisterUser(user *User) error
	Login(email string, password string) (*User, error)
}

type userService struct {
	repo UserRepository
}

//NewService Constructor
func NewService(repo UserRepository) Service {
	return &userService{
		repo: repo,
	}
}

func (u *userService) RegisterUser(user *User) error {
	u.repo.Create(user)
	return nil
}

func (u *userService) Login(email string, password string) (*User, error) {

	usr, err := u.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if usr.Password != password {
		return nil, errors.New("Invalid password")
	}

	return usr, nil
}
