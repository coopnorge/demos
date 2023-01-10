package grpcsvc

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	greeterpb "gitlab.com/aucampia/eg/service-golang/internal/generated/proto/example/greeter/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service services offers API server requests
type Service struct {
	greeterpb.UnimplementedGreeterAPIServer
}

// NewService creates a new Servicer object
func NewService() (servicer *Service, err error) {
	return &Service{}, nil
}

// Greet returns ...
func (s *Service) Greet(ctx context.Context, req *greeterpb.GreetRequest) (resp *greeterpb.GreetResponse, err error) {
	err = req.ValidateAll()
	if err != nil {
		logrus.WithError(err).Error("invalid request")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	logrus.WithField("name", req.Name).Info("validation passed, greeting")
	resp = &greeterpb.GreetResponse{Message: fmt.Sprintf("Hello %s", req.Name)}
	return resp, nil
}

var (
	_ greeterpb.GreeterAPIServer = (*Service)(nil)
)
