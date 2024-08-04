package database

const rowCountByName = `
SELECT count(*) FROM links
WHERE name = $1`

const insertLink = `
    INSERT INTO links
        (name, target, created_from)
    VALUES
        (:name, :target, :created_from)
`
