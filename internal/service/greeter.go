package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	validate "github.com/bufbuild/protovalidate-go"
	v1 "why.com/buf/api/helloworld/v1"
	"why.com/buf/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	log.Info(in.Name)
	validator, err := validate.New()
	if err != nil {
		return nil, err
	}
	if err := validator.Validate(in); err != nil {
		return nil, err
	}
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
