package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)

	lis,err :=net.Listen("tcp","9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis);err !=nil{
		log.Fatalf("failed to serve: %s",&err)
	}
	
}
