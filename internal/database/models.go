// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"time"
)

type User struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
