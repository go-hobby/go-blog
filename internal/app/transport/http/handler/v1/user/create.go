package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	errorsx "go-blog/internal/app/pkg/errors"
	"go-blog/internal/app/service/user"
	userGrpc "go-blog/internal/app/transport/grpc/api/scaffold/v1/user"
	"go-blog/internal/app/transport/http/pkg/response"
)

// Create 创建用户
// @Router       /v1/user [post]
// @Summary      创建用户
// @Description  创建用户
// @Tags         用户
// @Accept       json
// @Produce      json
// @Param        data  body      user.CreateRequest        true  "用户信息"  format(string)
// @Success      200   {object}  example.Success           "成功响应"
// @Failure      500   {object}  example.ServerError       "服务器出错"
// @Failure      400   {object}  example.ClientError       "客户端请求错误（code 类型应为 int，string 仅为了表达多个错误码）"
// @Failure      401   {object}  example.Unauthorized      "登陆失效"
// @Failure      403   {object}  example.PermissionDenied  "没有权限"
// @Failure      404   {object}  example.ResourceNotFound  "资源不存在"
// @Failure      429   {object}  example.TooManyRequest    "请求过于频繁"
// @Security     Authorization
func (h *Handler) Create(ctx *gin.Context) {
	req := new(user.CreateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil { //绑定参数
		h.logger.Error(err)
		response.Error(ctx, errorsx.ValidateError())
		return
	}
	reqCtx := ctx.Request.Context()

	conn, err := h.grpcClient.DialInsecure(reqCtx, h.conf.Services.Self) //grpc 监听
	if err != nil {
		h.logger.Error(err)
		response.Error(ctx, errorsx.ServerError())
		return
	}

	client := userGrpc.NewUserClient(conn)                                                            //链接grpc user
	_, err = client.Create(reqCtx, &userGrpc.CreateRequest{Name: req.Name, Age: 1, Phone: req.Phone}) //调用grpc Create
	if err != nil {
		e := errorsx.FromGRPCError(err)
		response.Error(ctx, fmt.Errorf("GRPC 调用错误：%s", e.Message))
		return
	}
	//	_, err := h.service.Create(ctx.Request.Context(), *req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx)
	return
}
