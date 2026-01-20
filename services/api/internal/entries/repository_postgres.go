package entries

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) Create(ctx context.Context, e StoredEntry) (StoredEntry, error) {
	analysisJSON, err := json.Marshal(e.Analysis)
	if err != nil {
		return StoredEntry{}, err
	}

	metadataJSON, err := json.Marshal(e.Metadata)
	if err != nil {
		return StoredEntry{}, err
	}

	const q = `
INSERT INTO entries (entry_id, created_at, text, source, metadata, analysis)
VALUES ($1, $2, $3, $4, $5::jsonb, $6::jsonb)
RETURNING entry_id, created_at, text, source, metadata, analysis
`

	var out StoredEntry
	var outMetadataBytes []byte
	var outAnalysisBytes []byte

	err = r.pool.QueryRow(ctx, q,
		e.EntryID,
		e.CreatedAt,
		e.Text,
		e.Source,
		metadataJSON,
		analysisJSON,
	).Scan(
		&out.EntryID,
		&out.CreatedAt,
		&out.Text,
		&out.Source,
		&outMetadataBytes,
		&outAnalysisBytes,
	)
	if err != nil {
		return StoredEntry{}, err
	}

	if len(outMetadataBytes) > 0 {
		_ = json.Unmarshal(outMetadataBytes, &out.Metadata)
	}
	if len(outAnalysisBytes) > 0 {
		_ = json.Unmarshal(outAnalysisBytes, &out.Analysis)
	}

	// Ensure UTC for consistency
	out.CreatedAt = out.CreatedAt.UTC()
	return out, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (StoredEntry, bool, error) {
	const q = `
SELECT entry_id, created_at, text, source, metadata, analysis
FROM entries
WHERE entry_id = $1
LIMIT 1
`

	var out StoredEntry
	var metadataBytes []byte
	var analysisBytes []byte

	err := r.pool.QueryRow(ctx, q, id).Scan(
		&out.EntryID,
		&out.CreatedAt,
		&out.Text,
		&out.Source,
		&metadataBytes,
		&analysisBytes,
	)
	if err != nil {
		// pgx returns an error on no rows - treat as not found
		// We avoid importing pgx just for pgx.ErrNoRows by checking string would be gross.
		// Better: use Query and check Next().
		rows, qerr := r.pool.Query(ctx, q, id)
		if qerr != nil {
			return StoredEntry{}, false, qerr
		}
		defer rows.Close()

		if !rows.Next() {
			return StoredEntry{}, false, nil
		}
		// If we got here, something else went wrong earlier.
		return StoredEntry{}, false, err
	}

	if len(metadataBytes) > 0 {
		_ = json.Unmarshal(metadataBytes, &out.Metadata)
	}
	if len(analysisBytes) > 0 {
		_ = json.Unmarshal(analysisBytes, &out.Analysis)
	}

	out.CreatedAt = out.CreatedAt.UTC()
	return out, true, nil
}
