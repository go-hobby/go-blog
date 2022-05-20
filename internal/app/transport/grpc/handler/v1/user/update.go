package user

import (
	"context"
	"go-blog/internal/app/service/user"
	pb "go-blog/internal/app/transport/grpc/api/scaffold/v1/user"
)

func (h *Handler) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	svcReq := user.UpdateRequest{
		Id:    req.Id,
		Name:  req.Name,
		Age:   int8(req.Age),
		Phone: req.Phone,
	}

	_, err := h.service.Update(ctx, svcReq)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{}, nil
}
