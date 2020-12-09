package service

import (
	"Go-000/Week02/dao"
)

func GetUserName(userID string) (string, error) {
	user, err := dao.GetUser(userID)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
