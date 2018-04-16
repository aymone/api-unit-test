package domain_test

import (
	"testing"

	"github.com/aymone/api-unit-test/domain"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	t.Run("empty id", func(t *testing.T) {
		u := new(domain.User)

		err := u.Validate()

		assert.Error(t, err)
		assert.Equal(t, domain.ErrEmptyID, err)
	})

	t.Run("zero id", func(t *testing.T) {
		u := new(domain.User)
		u.ID = "0"

		err := u.Validate()

		assert.Error(t, err)
		assert.Equal(t, domain.ErrZeroID, err)
	})

	t.Run("valid id", func(t *testing.T) {
		u := new(domain.User)
		u.ID = "123"

		err := u.Validate()

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}
