//+build wireinject

package main

import (
	service2 "Go-001/Week04/homework/internal/domain/user/service"
	"Go-001/Week04/homework/internal/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitHomeworkService(db *gorm.DB) service.HomeworkService {
	panic(wire.Build(service2.UserDomainWireSet, service.NewBasicHomeworkService))
}
