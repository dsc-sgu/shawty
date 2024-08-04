package database

const schema = `
CREATE TABLE IF NOT EXISTS links (
    name VARCHAR(256) PRIMARY KEY,
    target TEXT,
    ts TIMESTAMP WITH TIME ZONE DEFAULT now(),
    created_from VARCHAR(64) DEFAULT 'unknown'
)`
