package greet

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"go-blog/internal/app/service/greet"
)

type HandlerInterface interface {
	Hello(ctx *gin.Context)
}

var _ HandlerInterface = (*Handler)(nil)

type Handler struct {
	logger  *log.Helper
	service greet.ServiceInterface
}

func NewHandler(
	logger log.Logger,
	service greet.ServiceInterface,
) *Handler {
	return &Handler{
		logger:  log.NewHelper(logger),
		service: service,
	}
}
