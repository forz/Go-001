package client

import (
	pb "Go-001/Week04/apis/homework/pb"
	service "Go-001/Week04/homework/internal/service"
	endpoint1 "Go-001/Week04/homework/pkg/endpoint"
	"context"
	"errors"

	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.HomeworkService, error) {
	var getUserNameEndpoint endpoint.Endpoint
	{
		getUserNameEndpoint = grpc1.NewClient(conn, "pb.Homework", "GetUserName", encodeGetUserNameRequest, decodeGetUserNameReply, pb.GetUserNameReply{}, options["GetUserName"]...).Endpoint()
	}

	return endpoint1.Endpoints{GetUserNameEndpoint: getUserNameEndpoint}, nil
}

// encodeGetUserNameRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetUserName request to a gRPC request.
// TODO implement the encoder
func encodeGetUserNameRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Homework' Encoder is not impelemented")
}

// decodeGetUserNameReply is a transport/grpc.DecodeReplyFunc that converts
// a gRPC concat reply to a user-domain concat reply.
// TODO implement the decoder
func decodeGetUserNameReply(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Homework' Decoder is not impelemented")
}
