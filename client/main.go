package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mybiao/test/mess"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
)

func main() {
	reporter := zipkinhttp.NewReporter("http://abc.mybiao.top:9411/api/v2/spans")
	tracer, _ := zipkin.NewTracer(reporter)
	conn, err := grpc.Dial(":8000", grpc.WithInsecure(), grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracer)))
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	client := mess.NewUpperClient(conn)
	user := mess.User{
		Str: "this is client",
	}
	u, err := client.Hello(context.Background(), &user)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(u)
	fmt.Println("start call getCompany()")
	r := mess.Request{
		CompanyId: "aboisj=-32jnsdkj2jbjv",
	}
	gp := mess.NewCompanyClient(conn)
	res, err := gp.GetCompany(context.Background(), &r)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
	str := mess.NewMyStreamClient(conn)
	my, err := str.GetStream(context.Background(), &mess.ReqData{Data: "hello"})
	if err != nil {
		log.Println(err)
	}
	for {
		msg, err := my.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(msg)
	}

}
