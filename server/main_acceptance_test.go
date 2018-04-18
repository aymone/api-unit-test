// +build acceptance

package server_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aymone/api-unit-test/storage"
)

func TestMain(m *testing.M) {
	session, err := storage.NewSession()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	session.DB(storage.DBName).DropDatabase()

	code := m.Run()
	os.Exit(code)
}
