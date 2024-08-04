package db

import (
	"fmt"

	"github.com/dsc-sgu/atcc/internal/config"
	"github.com/jmoiron/sqlx"
)

var C *sqlx.DB

func Connect() (db *sqlx.DB, err error) {
	db, err = sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
			config.C.Pg.Username,
			config.C.Pg.Password,
			config.C.Pg.Host,
			config.C.Pg.Port,
			config.C.Pg.Name,
		),
	)
	return
}
