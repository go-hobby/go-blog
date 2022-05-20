package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"go-blog/internal/app/component/client/grpc"
	"go-blog/internal/app/component/trace"
	"go-blog/internal/app/config"
	"go-blog/internal/app/service/user"
)

type HandlerInterface interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Detail(ctx *gin.Context)
	List(ctx *gin.Context)
}

var _ HandlerInterface = (*Handler)(nil)

type Handler struct {
	logger     *log.Helper
	service    user.ServiceInterface
	trace      *trace.Tracer  //链路追踪
	grpcClient *grpc.Client   //grpc服务
	conf       *config.Config //配置
}

func NewHandler(
	logger log.Logger,
	service user.ServiceInterface,
	trace *trace.Tracer,
	grpcClient *grpc.Client,
	conf *config.Config,
) *Handler {
	return &Handler{
		logger:     log.NewHelper(logger),
		service:    service,
		trace:      trace,
		grpcClient: grpcClient,
		conf:       conf,
	}
}
