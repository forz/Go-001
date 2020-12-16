package server

import (
	pb "Go-001/Week04/apis/homework/pb"
	service "Go-001/Week04/homework/internal/service"
	endpoint "Go-001/Week04/homework/pkg/endpoint"
	transport "Go-001/Week04/homework/pkg/transport"

	endpoint1 "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
)

func New(h service.HomeworkService) pb.HomeworkServer {
	svc := service.New(getServiceMiddleware(), h)
	eps := endpoint.New(svc, getEndpointMiddleware())
	options := map[string][]grpc.ServerOption{}
	// Add your GRPC options here

	server := transport.New(eps, options)
	return server
}
func getServiceMiddleware() (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = addDefaultServiceMiddleware(mw)
	// Append your middleware here

	return
}
func getEndpointMiddleware() (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	// Add you endpoint middleware here

	return
}
