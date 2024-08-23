package database

const rowCountByName = `
SELECT count(*) FROM links
WHERE name = $1 AND NOT deleted`

const findByName = `
SELECT * FROM links
WHERE name = $1 AND NOT deleted
LIMIT 1`

const insertLink = `
INSERT INTO links (id, name, target, created_from)
VALUES (:id, :name, :target, :created_from)`
