-- Local development schema for Maximus.
-- Apply with: psql "$DSN" -f dev/setup.sql

CREATE TABLE IF NOT EXISTS documents (
    id      BIGSERIAL PRIMARY KEY,
    title   TEXT NOT NULL,
    body    TEXT NOT NULL,
    tags    TEXT,
    -- Operational metadata — should NOT trigger re-embedding.
    view_count    INTEGER NOT NULL DEFAULT 0,
    last_seen_at  TIMESTAMPTZ
);

-- Publication for Maximus to subscribe to.
-- Replication slot is created by Maximus on first run.
CREATE PUBLICATION maximus_pub FOR TABLE documents;
