package test

import (
	"log"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type fixtures struct {
	*testfixtures.Loader
	*gorm.DB
}

func newFixtures() (*fixtures, error) {
	var err error
	db, err := gorm.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}

	f, err := testfixtures.New(
		testfixtures.Database(db.DB()),            // You database connection
		testfixtures.Dialect("sqlite"),            // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("../test/fixture"), // the directory containing the YAML files
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		return nil, err
	}
	return &fixtures{f, db}, nil
}

func (f *fixtures) Prepare(t *testing.T) {
	if err := f.Load(); err != nil {
		log.Fatalf("cannot load fixtures, err: %+v", err)
	}
}
