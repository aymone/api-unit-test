package user_test

import (
	"testing"

	"github.com/aymone/api-unit-test/domain"
	"github.com/aymone/api-unit-test/domain/user"

	"github.com/stretchr/testify/assert"
)

type userStorageMock struct {
	GetFn      func(id string) (*domain.User, error)
	GetFnCount int

	InsertFn      func(user *domain.User) (*domain.User, error)
	InsertFnCount int
}

func (s *userStorageMock) Get(id string) (*domain.User, error) {
	s.GetFnCount++
	return s.GetFn(id)
}

func (s *userStorageMock) Insert(user *domain.User) (*domain.User, error) {
	s.InsertFnCount++
	return s.InsertFn(user)
}

func TestUserService(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		expectedId := "goRocks"
		userMock := &userStorageMock{
			GetFn: func(id string) (*domain.User, error) {
				return &domain.User{
					ID: expectedId,
				}, nil
			},
		}

		userService := user.NewService(userMock)

		u, err := userService.Get(expectedId)

		assert.NoError(t, err)
		assert.NotNil(t, u)

		expectedCount := 1
		assert.Equal(t, expectedCount, userMock.GetFnCount)
	})

	t.Run("insert", func(t *testing.T) {
		expectedUser := &domain.User{
			ID: "goRocks",
		}

		userMock := &userStorageMock{
			InsertFn: func(user *domain.User) (*domain.User, error) {
				return user, nil
			},
		}

		userService := user.NewService(userMock)

		u, err := userService.Insert(expectedUser)

		assert.NoError(t, err)
		assert.NotNil(t, u)

		expectedCount := 1
		assert.Equal(t, expectedCount, userMock.InsertFnCount)
	})
}
