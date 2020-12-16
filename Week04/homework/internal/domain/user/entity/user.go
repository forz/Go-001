package entity

import (
	"Go-001/Week04/homework/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("record not found")

type UserEntity interface {
	GetUserName(ctx context.Context, id string) (name string, err error)
}
type basicUserEntity struct {
	DB *gorm.DB
}

func NewBasicUserEntity(db *gorm.DB) UserEntity {
	return &basicUserEntity{DB: db}
}

func (u *basicUserEntity) GetUserName(ctx context.Context, id string) (name string, err error) {
	var user repository.User
	// err = u.DB.First(&user, id).Error
	// if err != nil {
	// 	// 屏蔽可预期的gorm错误,用entity的错误.
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		err = ErrRecordNotFound
	// 	}
	// 	return "", perrors.Wrap(err, fmt.Sprintf("id:%s", id))
	// }
	user, _ = mocksql(id)
	return user.Name, nil
}
func mocksql(id string) (user repository.User, err error) {
	return repository.User{
		ID:   "aa",
		Name: "girl",
		Age:  14,
	}, nil
}
