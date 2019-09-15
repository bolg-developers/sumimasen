package controllers

import (
	"context"
	"github.com/bolg-developers/sumimasen/server/app/services"
	"github.com/bolg-developers/sumimasen/server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SumimasenController struct {}

func (svc *SumimasenController) Sumimasen(ctx context.Context, req *proto.SumimasenRequest) (*proto.SumimasenResponse, error) {
	ret, err := services.GetSumimasenService().Sumimasen()
	if err != nil {
		return nil, status.New(codes.Internal, "sorry").Err()
	}
	return &proto.SumimasenResponse{Message: ret.Message}, nil
}
