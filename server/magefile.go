//go:build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	rdb "github.com/mrngsht/realworld-goa-react/myrdb"
)

func Setup() error {
	if err := sh.RunV("go", "install", "goa.design/goa/v3/cmd/goa@v3"); err != nil {
		return err
	}
	if err := sh.RunV("go", "install", "github.com/pressly/goose/v3/cmd/goose@v3.22.1"); err != nil {
		return err
	}
	if err := sh.RunV("go", "install", "github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0"); err != nil {
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
	rdbMigrationDirPath = "./myrdb/migrations"
	rdbSchemaFilePath   = "./myrdb/schema.sql"
)

var (
	gooseOpt = []string{"-dir", rdbMigrationDirPath, "postgres", rdb.LocalConnectionString}
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

func (Migration) Schema() error {
	// https://github.com/pressly/goose/issues/278#issuecomment-1921605230
	pgdump :=
		`pg_dump realworld \
  -h localhost \
  -U postgres \
  --schema-only \
  --no-comments \
  --quote-all-identifiers \
  -T public.goose_db_version \
  -T public.goose_db_version_id_seq | sed \
    -e '/^--.*/d' \
    -e '/^SET /d' \
    -e '/^[[:space:]]*$/d' \
    -e '/^SELECT pg_catalog./d' \
    -e '/^ALTER TABLE .* OWNER TO "postgres";/d' \
    -e 's/"public"\.//'`

	out, err := sh.Output("docker", "compose", "exec", "-T", "db", "sh", "-c", pgdump)
	if err != nil {
		return err
	}
	if err := os.WriteFile(rdbSchemaFilePath, []byte(out), 0644); err != nil {
		return err
	}
	return nil
}

type Sqlc mg.Namespace

const (
	rdbSqlcYamlPath = "myrdb/sqlc.yaml"
)

func (Sqlc) Gen() error {
	return sh.RunV("sqlc", "generate", "-f", rdbSqlcYamlPath)
}
