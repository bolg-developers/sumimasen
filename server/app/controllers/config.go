package controllers

import "github.com/bolg-developers/sumimasen/server/proto"

var (
	sumimasenController proto.SumimasenServer
)

func init() {
	sumimasenController = newSumimasenController()
}

func newSumimasenController() proto.SumimasenServer {
	return &SumimasenController{}
}

func GetSumimasenController() proto.SumimasenServer {
	return sumimasenController
}