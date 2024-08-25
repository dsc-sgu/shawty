package database

const rowCountByName = `
SELECT count(*) FROM links
WHERE name = $1 AND NOT deleted`

const findByName = `
SELECT * FROM link_visits
WHERE name = $1
LIMIT 1`

const insertLink = `
INSERT INTO links (id, name, target, created_from)
VALUES (:id, :name, :target, :created_from)`

const deleteLink = `
UPDATE links SET deleted = true, last_update = now()
WHERE id = $1`

const linksVisits = `
SELECT * FROM link_visits
ORDER BY last_update DESC
LIMIT $1
OFFSET $2`
