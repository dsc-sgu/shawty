package database

import (
	"context"
	"fmt"

	"github.com/dsc-sgu/atcc/internal/config"
	"github.com/dsc-sgu/atcc/internal/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var C *connection

type connection struct {
	db *sqlx.DB
}

// Connect to the database using credentials provided in the config.
func Connect() (c *connection, err error) {
	var db *sqlx.DB
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
	if err != nil {
		return
	}

	c = &connection{
		db: db,
	}
	return
}

// Initialize database schema if required tables do not exist.
func (c *connection) InitSchema(ctx context.Context) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, schema); err != nil {
		log.S.Errorw(
			"Database query has failed, performing rollback",
			"error", err,
		)
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// Closes database connection.
func (c *connection) Close() {
	c.db.Close()
}
