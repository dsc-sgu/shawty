package database

import (
	"context"
	"fmt"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var C *connection

type connection struct {
	db *sqlx.DB
}

// Connects to the database using credentials provided in the config.
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

// Initializes database schema, if required tables do not exist.
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

func (c *connection) IsNameTaken(
	ctx context.Context,
	name string,
) (bool, error) {
	var count []int
	if err := c.db.SelectContext(ctx, &count, rowCountByName, name); err != nil {
		log.S.Errorw("Database query has failed", "error", err)
		return false, err
	}
	return count[0] != 0, nil
}

// Finds link by its name. The second return value is the indicator,
// whether or not a link with this name exists in the database.
func (c *connection) FindLinkByName(ctx context.Context, name string) (
	ShortenedLink,
	bool,
	error,
) {
	var links []ShortenedLink
	if err := c.db.SelectContext(ctx, &links, findByName, name); err != nil {
		log.S.Errorw("Database query has failed", "error", err)
		return ShortenedLink{}, false, err
	}
	return links[0], len(links) != 0, nil
}

func (c *connection) SaveLink(ctx context.Context, l ShortenedLink) error {
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.NamedExecContext(ctx, insertLink, &l); err != nil {
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
