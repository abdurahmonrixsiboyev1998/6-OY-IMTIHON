// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package storage

import (
	"database/sql"
)

type User struct {
	ID           string
	Name         string
	Email        sql.NullString
	PasswordHash string
}
