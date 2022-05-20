package greet

import (
	"github.com/go-kratos/kratos/v2/log"
	"go-blog/internal/app/service/greet"
	pb "go-blog/internal/app/transport/grpc/api/scaffold/v1/greet"
)

var _ pb.GreetServer = (*Handler)(nil)

type Handler struct {
	pb.UnimplementedGreetServer
	logger  *log.Helper
	service *greet.Service
}

func NewHandler(
	logger log.Logger,
	service *greet.Service,
) *Handler {
	return &Handler{
		logger:  log.NewHelper(logger),
		service: service,
	}
}
