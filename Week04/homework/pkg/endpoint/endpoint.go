package endpoint

import (
	service "Go-001/Week04/homework/internal/service"
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GetUserNameRequest collects the request parameters for the GetUserName method.
type GetUserNameRequest struct {
	Id string `json:"id"`
}

// GetUserNameReply collects the reply parameters for the GetUserName method.
type GetUserNameReply struct {
	Name string `json:"name"`
	Err  error  `json:"err"`
}

// MakeGetUserNameEndpoint returns an endpoint that invokes GetUserName on the service.
func MakeGetUserNameEndpoint(s service.HomeworkService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserNameRequest)
		name, err := s.GetUserName(ctx, req.Id)
		reply := GetUserNameReply{
			Err:  err,
			Name: name,
		}
		return reply, reply.Failed()
	}
}

// Failed implements Failer.
func (r GetUserNameReply) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by reply types.
// Reply encoders can check if replys are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetUserName implements Service. Primarily useful in a client.
func (e Endpoints) GetUserName(ctx context.Context, id string) (name string, err error) {
	request := GetUserNameRequest{Id: id}
	reply, err := e.GetUserNameEndpoint(ctx, request)
	if err != nil {
		return
	}
	return reply.(GetUserNameReply).Name, reply.(GetUserNameReply).Err
}
