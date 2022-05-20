package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/google/wire"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-blog/internal/app/config"
	"go-blog/internal/app/pkg/errors"
	"go-blog/internal/app/transport/http/api/docs"
	"go-blog/internal/app/transport/http/handler/v1/greet"
	"go-blog/internal/app/transport/http/handler/v1/trace"
	"go-blog/internal/app/transport/http/handler/v1/user"
	"go-blog/internal/app/transport/http/middleware/recover"
	"go-blog/internal/app/transport/http/pkg/response"
	"go-blog/internal/app/transport/http/pkg/swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var ProviderSet = wire.NewSet(
	New,
)

// New 返回 gin 路由对象
func New(
	loggerWriter *rotatelogs.RotateLogs,
	zLogger *zap.Logger,
	logger log.Logger,
	appConf *config.App,
	httpConf *config.HTTP,
	jwtConf *config.JWT,
	enforcer *casbin.Enforcer,
	greetHandler greet.HandlerInterface,
	traceHandler trace.HandlerInterface,
	userHandler user.HandlerInterface,
) *gin.Engine {
	if httpConf == nil {
		return nil
	}

	var output io.Writer
	if loggerWriter == nil {
		output = os.Stdout
	} else {
		output = io.MultiWriter(os.Stdout, loggerWriter)
	}
	gin.DefaultWriter = output
	gin.DefaultErrorWriter = output
	gin.DisableConsoleColor()

	switch appConf.Env {
	case config.Local:
		gin.SetMode(gin.DebugMode)
	case config.Test:
		gin.SetMode(gin.TestMode)
	case config.Prod:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(ginzap.Ginzap(zLogger.WithOptions(zap.AddCallerSkip(4)), time.RFC3339, false))

	router.Use(recover.CustomRecoveryWithZap(zLogger.WithOptions(zap.AddCallerSkip(4)), true, func(c *gin.Context, err interface{}) {
		response.Error(c, errors.ServerError())
		c.Abort()
	}))
	router.Use(otelgin.Middleware(appConf.Name))

	rg := router.Group("/")
	extAddrSubs := strings.SplitN(httpConf.ExternalAddr, "/", 2)
	if len(extAddrSubs) == 2 {
		rg = router.Group("/" + extAddrSubs[1])
	}

	rg.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong"); return })
	// 注册 api 路由组
	apiGroup := rg.Group("/api")
	{
		apiGroup.Use(cors.Default(), Cors()) // 允许跨越
		/*if jwtConf != nil {
			if jwtConf.Key != "" {
				apiGroup.Use(jwtmd.New(
					jwtConf.Key,
					jwtmd.WithLogger(log.NewHelper(logger)),
					jwtmd.WithErrorResponseBody(response.NewBody(int(errors.ServerErrorCode), errors.ServerErrorCode.String(), nil)),
					jwtmd.WithValidateFailedResponseBody(response.NewBody(int(errors.UnauthorizedCode), errors.UnauthorizedCode.String(), nil)),
				).Validate())
			}
		}*/
		/*
			if enforcer != nil {
				apiGroup.Use(casbinmd.New(
					enforcer,
					func(ctx *gin.Context) ([]interface{}, error) {
						// TODO
						return nil, nil
					},
					casbinmd.WithLogger(log.NewHelper(logger)),
					casbinmd.WithErrorResponseBody(response.NewBody(int(errors.ServerErrorCode), errors.ServerErrorCode.String(), nil)),
					casbinmd.WithValidateFailedResponseBody(response.NewBody(int(errors.UnauthorizedCode), errors.UnauthorizedCode.String(), nil)),
				).Validate())
			}
		*/

		// swagger 配置
		if appConf.Env == config.Local {
			docs.SwaggerInfo.Host = httpConf.Addr
			if len(extAddrSubs) > 0 {
				docs.SwaggerInfo.Host = extAddrSubs[0]
			}
			docs.SwaggerInfo.BasePath = apiGroup.BasePath()

			swagger.Setup(router, swagger.Config{
				Path: apiGroup.BasePath() + "/docs",
				Option: func(c *ginSwagger.Config) {
					c.DefaultModelsExpandDepth = -1
				},
			})
		}

		apiV1Group := apiGroup.Group("/v1")
		{
			// TODO 编写路由

			apiV1Group.GET("/greet", greetHandler.Hello)
			apiV1Group.GET("/trace", traceHandler.Example)

			apiV1Group.GET("/users", userHandler.List)
			apiV1Group.GET("/user/:id", userHandler.Detail)
			apiV1Group.POST("/user", userHandler.Create)
			apiV1Group.PUT("/user/:id", userHandler.Update)
			apiV1Group.DELETE("/user/:id", userHandler.Delete)
		}
	}

	return router
}

func Cors() gin.HandlerFunc {
	var ctx *gin.Context
	traceID := tracing.TraceID()(ctx).(string)
	return func(c *gin.Context) {
		c.Header("token", traceID)
		c.Next()
	}
}
