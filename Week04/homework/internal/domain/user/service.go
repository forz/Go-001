package user

import (
	"Go-001/Week04/homework/internal/domain/user/entity"
	"context"

	"github.com/google/wire"
)

var UserDomainWireSet = wire.NewSet(entity.NewBasicUserEntity, NewBasicDomainService)

type DomainService interface {
	// Add your methods here
	GetUserName(ctx context.Context, id string) (name string, err error)
}

type basicDomainService struct {
	entity entity.UserEntity
}

// NewBasicHomeworkService returns a naive, stateless implementation of HomeworkService.
func NewBasicDomainService(e entity.UserEntity) DomainService {
	return &basicDomainService{entity: e}
}

func (u *basicDomainService) GetUserName(ctx context.Context, id string) (name string, err error) {
	return u.entity.GetUserName(ctx, id)
}
