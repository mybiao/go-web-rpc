package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mybiao/test/mess"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	mess.RegisterUpperServer(server, &mess.ImplHelloUpperServer{})
	mess.RegisterCompanyServer(server, &mess.ImplHelloUpperServer{})
	mess.RegisterMyStreamServer(server, &mess.ImplHelloUpperServer{})
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
	}
	server.Serve(lis)

}
func format() string {
	now := time.Now()
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}
