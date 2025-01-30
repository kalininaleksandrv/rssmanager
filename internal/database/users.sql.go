// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package database

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(id, name, created_at, updated_at) 
VALUES ($1, $2, $3, $4) 
RETURNING id, name, created_at, updated_at
`

type CreateUserParams struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const selectUser = `-- name: SelectUser :one
SELECT id, name, created_at, updated_at FROM users
`

func (q *Queries) SelectUser(ctx context.Context) (User, error) {
	row := q.db.QueryRowContext(ctx, selectUser)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
