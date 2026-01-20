CREATE TABLE IF NOT EXISTS entries (
  entry_id   TEXT PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL,
  text       TEXT NOT NULL,
  source     TEXT NOT NULL DEFAULT '',
  metadata   JSONB NOT NULL DEFAULT '{}'::jsonb,
  analysis   JSONB NOT NULL DEFAULT '{}'::jsonb
);

CREATE INDEX IF NOT EXISTS idx_entries_created_at ON entries (created_at DESC);
