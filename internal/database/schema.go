package database

const schema = `
CREATE TABLE IF NOT EXISTS links (
    id UUID PRIMARY KEY,
    name VARCHAR(256),
    target TEXT,
    deleted BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    created_from VARCHAR(64) DEFAULT 'unknown',
    last_update TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS visits (
    id UUID PRIMARY KEY,
    link_id UUID,
    tag VARCHAR(256) DEFAULT '',
    host VARCHAR(256) DEFAULT '',
    "timestamp" TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE OR REPLACE VIEW link_visits AS (
	WITH visit_count AS (
		SELECT v.link_id, count(*) total_count FROM visits v
		GROUP BY (v.link_id)
	)
	SELECT
		l.name,
		l.target,
		l.created_at,
		l.created_from,
		l.last_update,
		coalesce(vc.total_count, 0) total_visits
	FROM links l
	LEFT JOIN visit_count vc ON l.id = vc.link_id
);`
