-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $2, updated_at = $3
WHERE id = $1
RETURNING *;