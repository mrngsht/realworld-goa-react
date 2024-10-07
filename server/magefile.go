//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Setup() error {
	if err := sh.RunV("go", "install", "goa.design/goa/v3/cmd/goa@v3"); err != nil {
		return err
	}
	if err := sh.RunV("go", "install", "github.com/pressly/goose/v3/cmd/goose@v3.22.1"); err != nil {
		return err
	}
	return nil
}

type Goa mg.Namespace

func (Goa) Gen() error {
	return sh.RunV("goa", "gen", "github.com/mrngsht/realworld-goa-react/design")
}

type Migration mg.Namespace

const (
	rdbConnectionString = "host=localhost user=postgres password=postgres dbname=realworld sslmode=disable"
	rdbMigrationDirPath = "./rdb/migrations"
)

var (
	gooseOpt = []string{"-dir", rdbMigrationDirPath, "postgres", rdbConnectionString}
)

func (Migration) New(name string) error {
	return sh.RunV("goose", append(gooseOpt, "create", name, "sql")...)
}

func (Migration) Up() error {
	return sh.RunV("goose", append(gooseOpt, "up")...)
}

func (Migration) Down() error {
	return sh.RunV("goose", append(gooseOpt, "down")...)
}

func (Migration) Status() error {
	return sh.RunV("goose", append(gooseOpt, "status")...)
}
