// +build acceptance

package storage_test

import (
	"testing"

	"github.com/aymone/api-unit-test/domain"
	"github.com/aymone/api-unit-test/storage"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	session, err := storage.NewSession()
	assert.NoError(t, err)

	storage, err := storage.NewUserStorage(session)
	assert.NoError(t, err)
	assert.NotNil(t, storage)

	userID := "01"

	t.Run("insert", func(t *testing.T) {
		user := &domain.User{
			ID: userID,
		}

		user, err := storage.Insert(user)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userID, user.ID)
	})

	t.Run("get", func(t *testing.T) {
		user := &domain.User{}
		user, err := storage.Get(userID)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userID, user.ID)
	})
}
