package user

import (
	"github.com/aymone/api-unit-test/domain"
)

type userStorage interface {
	Get(id string) (*domain.User, error)
	Insert(user *domain.User) (*domain.User, error)
}

type userService struct {
	storage userStorage
}

// NewService returs a user service
func NewService(s userStorage) *userService {
	return &userService{
		storage: s,
	}
}

// Get retrieve user by ID
func (u *userService) Get(id string) (*domain.User, error) {
	return u.storage.Get(id)
}

// Insert register an user
func (u *userService) Insert(user *domain.User) (*domain.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return u.storage.Insert(user)
}
