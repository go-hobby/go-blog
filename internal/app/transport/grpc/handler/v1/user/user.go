package user

import (
	"github.com/go-kratos/kratos/v2/log"
	userrepo "go-blog/internal/app/repository/user"
	usersvc "go-blog/internal/app/service/user"
	pb "go-blog/internal/app/transport/grpc/api/scaffold/v1/user"
)

var _ pb.UserServer = (*Handler)(nil)

type Handler struct {
	pb.UnimplementedUserServer
	logger  *log.Helper                  //日志
	service *usersvc.Service             //业务层
	repo    userrepo.RepositoryInterface // dao
}

func NewHandler(
	logger log.Logger,
	service *usersvc.Service,
	repo userrepo.RepositoryInterface,
) *Handler {
	return &Handler{
		logger:  log.NewHelper(logger),
		service: service,
		repo:    repo,
	}
}
