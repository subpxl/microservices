package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "loggerService/grpcproto"

	"github.com/golang/protobuf/proto"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RedisSetup() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// ctx := context.Background()
	// err := client.Set(ctx, "foo", "bar", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// val, err := client.Get(ctx, "foo").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("foo", val)
	return client

}

var Client *redis.Client

func main() {
	Client = RedisSetup()
	grpcServer := grpc.NewServer()
	server := &messageServer{}
	pb.RegisterLoggerServiceServer(grpcServer, server) // Register the gRPC service
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println("starting grpc server on 5000")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("failed t start", err)
	}
}

type messageServer struct {
	pb.UnimplementedLoggerServiceServer
}

func (s *messageServer) SayHello(ctx context.Context, req *pb.Message) (*pb.Message, error) {
	fmt.Println("req is ", req)
	ctxb := context.Background()
	binaryMessage, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	err = Client.Set(ctxb, "foo", binaryMessage, 0).Err()
	if err != nil {
		panic(err)
	}

	storedBinaryMessage, err := Client.Get(ctxb, "foo").Result()
	if err != nil {
		panic(err)
	}
	// result := "added log"

	storedMessage := &pb.Message{}
	err = proto.Unmarshal([]byte(storedBinaryMessage), storedMessage)
	if err != nil {
		panic(err)
	}

	fmt.Print(storedMessage)

	// result = string( storedMessage)

	return storedMessage, nil
	// return &pb.Message{Body: result}, nil
}
