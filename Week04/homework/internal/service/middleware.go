package service

import (
	"context"
	"log"
)

// Middleware describes a service middleware.
type Middleware func(HomeworkService) HomeworkService

type loggingMiddleware struct {
	next HomeworkService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a HomeworkService Middleware.
func LoggingMiddleware() Middleware {
	return func(next HomeworkService) HomeworkService {
		return &loggingMiddleware{next}
	}

}

func (l loggingMiddleware) GetUserName(ctx context.Context, id string) (name string, err error) {
	defer func() {
		log.Println("id", id, "name", name, "err", err, "method", "GetUserName")
	}()
	return l.next.GetUserName(ctx, id)
}
