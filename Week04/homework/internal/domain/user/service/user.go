package service

import (
	"Go-001/Week04/homework/internal/domain/user/entity"
	"context"

	"github.com/google/wire"
)

var UserDomainWireSet = wire.NewSet(entity.NewBasicUserEntity, NewBasicUserDomainService)

type UserDomainService interface {
	// Add your methods here
	GetUserName(ctx context.Context, id string) (name string, err error)
}

type basicUserDomainService struct {
	entity entity.UserEntity
}

// NewBasicHomeworkService returns a naive, stateless implementation of HomeworkService.
func NewBasicUserDomainService(e entity.UserEntity) UserDomainService {
	return &basicUserDomainService{entity: e}
}

func (u *basicUserDomainService) GetUserName(ctx context.Context, id string) (name string, err error) {
	return u.entity.GetUserName(ctx, id)
}
