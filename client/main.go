package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mybiao/test/mess"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
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
}
