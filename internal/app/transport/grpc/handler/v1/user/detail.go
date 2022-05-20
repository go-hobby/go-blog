package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go-blog/internal/app/service/user"
	pb "go-blog/internal/app/transport/grpc/api/scaffold/v1/user"
)

func (h *Handler) Detail(ctx context.Context, req *pb.DetailRequest) (*pb.DetailResponse, error) {
	svcReq := user.DetailRequest{
		Id: req.Id,
	}

	ret, err := h.service.Detail(ctx, svcReq)
	if err != nil {
		return nil, err
	}
	traceID := tracing.TraceID()(ctx).(string)
	resp := &pb.DetailResponse{
		Id:    ret.Id,
		Name:  traceID,
		Age:   int32(ret.Age),
		Phone: ret.Phone,
	}

	return resp, nil
}
