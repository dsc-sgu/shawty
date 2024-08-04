package database

type ShortenedLink struct {
	Name        string `db:"name"`
	Target      string `db:"target"`
	CreatedFrom string `db:"created_from"`
}
