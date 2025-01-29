// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type Note struct {
	ID        int64
	UserID    int64
	Content   string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type User struct {
	ID       int64
	Username string
	Password string
	Roles    []string
	Active   bool
}
