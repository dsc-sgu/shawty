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
);`
