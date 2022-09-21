package service

import (
	"context"

	v1 "lusun/api/helloworld/v1"
	"lusun/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc             *biz.GreeterUsecase
	articleUseCase *biz.ArticleUseCase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, articleUseCase *biz.ArticleUseCase) *GreeterService {
	return &GreeterService{
		uc:             uc,
		articleUseCase: articleUseCase,
	}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}

	a, err := s.articleUseCase.Detail(1)
	println(a.LogID)

	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
