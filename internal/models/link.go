package models

import (
	"time"

	"github.com/google/uuid"
)

type Link struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Target      string    `db:"target"`
	Deleted     bool      `db:"deleted"`
	CreatedAt   time.Time `db:"created_at"`
	CreatedFrom string    `db:"created_from"`
	LastUpdate  time.Time `db:"last_update"`
}

type LinkWithVisits struct {
	Id          uuid.UUID `db:"id"           json:"id"`
	Name        string    `db:"name"         json:"name"`
	Target      string    `db:"target"       json:"target"`
	CreatedAt   time.Time `db:"created_at"   json:"created_at"`
	CreatedFrom string    `db:"created_from" json:"created_from"`
	LastUpdate  time.Time `db:"last_update"  json:"last_update"`
	TotalVisits int       `db:"total_visits" json:"total_visits"`
}

type Visit struct {
	Id        uuid.UUID `db:"id"        json:"id"`
	LinkId    uuid.UUID `db:"link_id"   json:"link_id"`
	Tag       string    `db:"tag"       json:"tag"`
	Host      string    `db:"host"      json:"host"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`
}
