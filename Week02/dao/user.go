// Package dao ...
package dao

import (
	"database/sql"

	"errors"

	perrors "github.com/pkg/errors"
)

// ErrRecordNotFound sentinel error not use pkg/errors
var ErrRecordNotFound = errors.New("record not found")

// User is user basic info.
type User struct {
	ID   string
	Name string
}

// mockSQL mock sql for return sql.ErrNoRows.
func mockSQL(userID string, user *User) error {
	return sql.ErrNoRows
}

// GetUser return a user info by userID.
func GetUser(userID string) (User, error) {
	var user User
	err := mockSQL(userID, &user)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrRecordNotFound
		}
		return user, perrors.Wrapf(err, "userID: %s", userID)
	}

	return user, nil
}
