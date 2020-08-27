package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mybiao/test/mess"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
)

func main() {
	reporter := zipkinhttp.NewReporter("http://abc.mybiao.top:9411/api/v2/spans")
	tracer, _ := zipkin.NewTracer(reporter)
	server := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)))
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
