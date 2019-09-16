package controllers

import (
	pb "github.com/bolg-developers/sumimasen/api"
)

var (
	sumimasenController pb.SumimasenServer
)

func init() {
	sumimasenController = newSumimasenController()
}

func newSumimasenController() pb.SumimasenServer {
	return &SumimasenController{}
}

func GetSumimasenController() pb.SumimasenServer {
	return sumimasenController
}
