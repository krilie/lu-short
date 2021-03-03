package dbmigrate

import (
	"database/sql"
	"embed"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/johejo/golang-migrate-extra/source/iofs"
)

// 相对目录
//go:embed migrations/*.sql
var sqlFiles embed.FS

func Migrate(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}

	d, err := iofs.New(sqlFiles, "migrations")
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		d, // 相对目录
		"mysql", driver)
	if err != nil {
		panic(err)
	}

	err = m.Migrate(210226) // <0 down,=0 noop,>0 up
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}
