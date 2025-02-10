-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFeedsByUserId :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: GetFeedsForFetchUpdate :many
SELECT * FROM feeds WHERE last_fetched_at IS NULL OR last_fetched_at < $1;

-- name: UpdateFeedLastFetch :one
UPDATE feeds
SET last_fetched_at = $2
WHERE id = $1
RETURNING *;