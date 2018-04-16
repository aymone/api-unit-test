package server

import (
	"github.com/aymone/api-unit-test/domain"
)

type userService interface {
	Get(id string) (*domain.User, error)
	Insert(user *domain.User) (*domain.User, error)
}

type Handler struct {
	userService userService
}

func NewHandler(u userService) *Handler {
	return &Handler{
		userService: u,
	}
}
