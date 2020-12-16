package transport

import (
	pb "Go-001/Week04/apis/homework/pb"
	endpoint "Go-001/Week04/homework/pkg/endpoint"
	"context"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeGetUserNameHandler creates the handler logic
func makeGetUserNameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserNameEndpoint, decodeGetUserNameRequest, encodeGetUserNameReply, options...)
}

// decodeGetUserNameRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserName request.
func decodeGetUserNameRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetUserNameRequest)
	return endpoint.GetUserNameRequest{
		Id: req.Id,
	}, nil
}

// encodeGetUserNameReply is a transport/grpc.EncodeReplyFunc that converts
// a user-domain reply to a gRPC reply.
func encodeGetUserNameReply(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.GetUserNameReply)
	return &pb.GetUserNameReply{
		Name: res.Name,
	}, nil
}
func (g *grpcServer) GetUserName(ctx context1.Context, req *pb.GetUserNameRequest) (*pb.GetUserNameReply, error) {
	_, rep, err := g.getUserName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserNameReply), nil
}
