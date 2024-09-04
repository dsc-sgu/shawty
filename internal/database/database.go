package database

import (
	"context"
	"fmt"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/log"
	"github.com/dsc-sgu/shawty/internal/models"
	"github.com/google/uuid"
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
	models.LinkWithVisits,
	bool,
	error,
) {
	var links []models.LinkWithVisits
	if err := c.db.SelectContext(ctx, &links, findByName, name); err != nil {
		log.S.Errorw("Database query has failed", "error", err)
		return models.LinkWithVisits{}, false, err
	}
	if len(links) != 0 {
		return links[0], true, nil
	} else {
		return models.LinkWithVisits{}, false, nil
	}
}

// Finds link by its ID. The second return value is the indicator,
// whether or not a link with this ID exists in the database.
func (c *connection) FindLinkById(ctx context.Context, id uuid.UUID) (
	models.LinkWithVisits,
	bool,
	error,
) {
	var links []models.LinkWithVisits
	if err := c.db.SelectContext(ctx, &links, findById, id); err != nil {
		log.S.Errorw("Database query has failed", "error", err)
		return models.LinkWithVisits{}, false, err
	}
	if len(links) != 0 {
		return links[0], true, nil
	} else {
		return models.LinkWithVisits{}, false, nil
	}
}

// Inserts link in the database.
func (c *connection) SaveLink(ctx context.Context, l models.Link) error {
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
		_ = tx.Rollback()
		return err
	}
	return nil
}

// Marks link as `deleted` in the database.
func (c *connection) DeleteLink(ctx context.Context, id uuid.UUID) error {
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, deleteLink, id); err != nil {
		log.S.Errorw(
			"Database query has failed, performing rollback",
			"error", err,
		)
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}
	return nil
}

// Queries records from the view, that joins `links` and `visits`
// tables, to count each link's visits.
func (c *connection) GetLinksWithVisits(
	ctx context.Context,
	page int,
	size int,
) (
	[]models.LinkWithVisits,
	error,
) {
	var links []models.LinkWithVisits
	if err := c.db.SelectContext(ctx, &links, linksVisits, size, page*size); err != nil {
		log.S.Errorw("Database query has failed", "error", err)
		return []models.LinkWithVisits{}, err
	}
	return links, nil
}

// Inserts visit in the database.
func (c *connection) SaveVisit(ctx context.Context, v models.Visit) error {
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.NamedExecContext(ctx, insertVisit, &v); err != nil {
		log.S.Errorw(
			"Database query has failed, performing rollback",
			"error", err,
		)
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}
	return nil
}

// Closes database connection.
func (c *connection) Close() {
	c.db.Close()
}
