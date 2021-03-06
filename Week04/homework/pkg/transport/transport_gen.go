// THIS FILE IS AUTO GENERATED BY Kepler DO NOT EDIT!!
package transport

import (
	pb "Go-001/Week04/apis/homework/pb"
	endpoint "Go-001/Week04/homework/pkg/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
)

// New makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	getUserName grpc.Handler
}

func New(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.HomeworkServer {
	return &grpcServer{getUserName: makeGetUserNameHandler(endpoints, options["GetUserName"])}
}
