//+build wireinject

package main

import (
	"Go-001/Week04/homework/internal/domain/user"
	"Go-001/Week04/homework/internal/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitHomeworkService(db *gorm.DB) service.HomeworkService {
	panic(wire.Build(user.UserDomainWireSet, service.NewBasicHomeworkService))
}
