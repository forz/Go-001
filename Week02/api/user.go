// Package api ...
package api

import (
	"Go-000/Week02/service"
)

// SameUserName compare 2 user's name.
func SameUserName(userIDA, userIDB string) (bool, error) {
	userNameA, err := service.GetUserName(userIDA)
	if err != nil {
		return false, err
	}

	userNameB, err := service.GetUserName(userIDB)
	if err != nil {
		return false, err
	}

	return userNameA == userNameB, nil
}
