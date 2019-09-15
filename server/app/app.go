package main

import (
	"log"
	"net"

	"github.com/bolg-developers/sumimasen/server/app/controllers"
	"github.com/bolg-developers/sumimasen/server/config"
	"github.com/bolg-developers/sumimasen/server/proto"
	"google.golang.org/grpc"
)

func main() {
	port, err := net.Listen("tcp", config.Env().ServerPort)
	if err != nil {
		log.Fatal("port listen error:", port)
	}

	srv := grpc.NewServer()

	proto.RegisterSumimasenServer(srv, controllers.GetSumimasenController())

	log.Fatal(srv.Serve(port))
}
