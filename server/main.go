package main

import (
	"log"
	"net"

	"github.com/mybiao/test/mess"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	mess.RegisterUpperServer(server, &mess.ImplHelloUpperServer{})
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
	}
	server.Serve(lis)
}
