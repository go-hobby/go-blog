package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go-blog/internal/app/pkg/errors"
	"net/http"
)

// Response HTTP 请求响应
func Response(ctx *gin.Context, httpStatusCode int, errorCode int, msg string, ops ...Option) {
	b := NewBody(errorCode, msg, nil)
	for _, op := range ops {
		op(b)
	}
	ctx.JSON(httpStatusCode, b)
}

// Success HTTP 成功响应
func Success(ctx *gin.Context, ops ...Option) {
	var c = ctx.Request.Context()
	traceID := tracing.TraceID()(c).(string)
	ops = append(ops, WithTraceID(traceID))
	Response(ctx, http.StatusOK, int(errors.SuccessCode), errors.SuccessCode.String(), ops...)
}

// Error HTTP 错误响应
func Error(ctx *gin.Context, err error, ops ...Option) {
	ae, ok := err.(*errors.Error)
	if !ok {
		ae = errors.ServerError(errors.WithMessage(err.Error()))
	}
	var c = ctx.Request.Context()
	traceID := tracing.TraceID()(c).(string)
	ops = append(ops, WithTraceID(traceID))
	Response(ctx, ae.Code.HTTPStatusCode(), int(ae.Code), ae.Message, ops...)
}
