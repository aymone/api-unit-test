// +build acceptance

package storage_test

import (
	"os"
	"testing"

	"github.com/aymone/api-unit-test/storage"
)

func TestMain(m *testing.M) {
	session, _ := storage.NewSession()
	session.DB(storage.DBName).DropDatabase()

	code := m.Run()
	os.Exit(code)
}
