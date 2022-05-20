package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"go-blog/internal/app/config"
	"go-blog/internal/app/transport/http/handler"
	"go-blog/internal/app/transport/http/router"
	"time"
)

var ProviderSet = wire.NewSet(
	handler.ProviderSet, //绑定
	router.ProviderSet,  //路由
	NewServer,
)

// NewServer 创建 HTTP 服务器
func NewServer(
	logger log.Logger,
	httpConf *config.HTTP,
	router *gin.Engine,
) *khttp.Server {
	if router == nil {
		return nil
	}

	var opts = []khttp.ServerOption{
		khttp.Logger(logger),
	}

	if httpConf.Network != "" {
		opts = append(opts, khttp.Network(httpConf.Network))
	}

	if httpConf.Addr != "" {
		opts = append(opts, khttp.Address(httpConf.Addr))
	}

	if httpConf.Timeout != 0 {
		opts = append(opts, khttp.Timeout(time.Duration(httpConf.Timeout)*time.Second))
	}

	srv := khttp.NewServer(opts...)
	srv.HandlePrefix("/", router)

	return srv
}
