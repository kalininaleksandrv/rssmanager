-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFeedsByUserId :many
SELECT * FROM feeds WHERE user_id = $1;