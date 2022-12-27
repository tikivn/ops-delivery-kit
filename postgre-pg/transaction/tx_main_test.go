//go:build integration_test
// +build integration_test

package transaction

import (
	"os"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"

	postgre_pg "github.com/tikivn/ops-delivery-kit/postgre-pg"
)

var (
	db *pg.DB
)

func TestMain(m *testing.M) {
	logrus.Infof("TestMain run")
	pgDB, cleanup, err := postgre_pg.SetupDBTest()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	db = pgDB

	code := m.Run()
	os.Exit(code)
}
