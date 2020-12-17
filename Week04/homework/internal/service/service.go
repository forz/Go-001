package service

import (
	"Go-001/Week04/homework/internal/domain/user"
	"context"
)

// HomeworkService describes the service.
type HomeworkService interface {
	// Add your methods here
	GetUserName(ctx context.Context, id string) (name string, err error)
}

type basicHomeworkService struct {
	svc user.DomainService
}

// NewBasicHomeworkService returns a naive, stateless implementation of HomeworkService.
func NewBasicHomeworkService(u user.DomainService) HomeworkService {
	return &basicHomeworkService{svc: u}
}

func (b *basicHomeworkService) GetUserName(ctx context.Context, id string) (name string, err error) {
	return b.svc.GetUserName(ctx, id)
}

// New returns a HomeworkService with all of the expected middleware wired in.
func New(middleware []Middleware, homeservice HomeworkService) HomeworkService {
	for _, m := range middleware {
		homeservice = m(homeservice)
	}
	return homeservice
}
