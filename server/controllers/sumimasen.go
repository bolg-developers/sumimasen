package controllers

import (
	"context"

	pb "github.com/bolg-developers/sumimasen/api"
	"github.com/bolg-developers/sumimasen/server/services"
)

type SumimasenController struct{}

func (ctl *SumimasenController) Sumimasen(ctx context.Context, req *pb.SumimasenRequest) (*pb.SumimasenResponse, error) {
	res := services.GetSumimaseService().Sumimasen()
	ret := &pb.SumimasenResponse{
		Message: res.Message,
	}
	return ret, nil
}
