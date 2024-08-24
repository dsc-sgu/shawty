package database

import (
	"time"

	"github.com/google/uuid"
)

type ShortenedLink struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Target      string    `db:"target"`
	Deleted     bool      `db:"deleted"`
	CreatedAt   time.Time `db:"created_at"`
	CreatedFrom string    `db:"created_from"`
	LastUpdate  time.Time `db:"last_update"`
}

type LinkVisits struct {
	Name        string    `db:"name"`
	Target      string    `db:"target"`
	CreatedAt   time.Time `db:"created_at"`
	CreatedFrom string    `db:"created_from"`
	LastUpdate  time.Time `db:"last_update"`
	TotalVisits int       `db:"total_visits"`
}

type Visit struct {
	Id        uuid.UUID `db:"id"`
	LinkId    uuid.UUID `db:"link_id"`
	Tag       string    `db:"tag"`
	Host      string    `db:"host"`
	Timestamp time.Time `db:"timestamp"`
}
