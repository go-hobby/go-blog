package trace

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"go-blog/internal/app/component/client/grpc"
	"go-blog/internal/app/component/trace"
	"go-blog/internal/app/config"
)

type HandlerInterface interface {
	Example(ctx *gin.Context)
}

var _ HandlerInterface = (*Handler)(nil)

type Handler struct {
	logger     *log.Helper
	conf       *config.Config
	trace      *trace.Tracer
	grpcClient *grpc.Client
}

func NewHandler(
	logger log.Logger,
	conf *config.Config,
	trace *trace.Tracer,
	grpcClient *grpc.Client,
) *Handler {
	return &Handler{
		logger:     log.NewHelper(logger),
		conf:       conf,
		trace:      trace,
		grpcClient: grpcClient,
	}
}
