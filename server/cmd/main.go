package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/bolg-developers/sumimasen/api"
	"github.com/bolg-developers/sumimasen/server/config"
	"github.com/bolg-developers/sumimasen/server/controllers"
	"google.golang.org/grpc"
)

func main() {
	fmt.Print("hello, world")
	port, err := net.Listen("tcp", ":"+config.Env().ServerPort)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	srv := grpc.NewServer()
	pb.RegisterSumimasenServer(srv, controllers.GetSumimasenController())

	log.Fatal(srv.Serve(port))
}
