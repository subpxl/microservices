package grpc

import (
	"context"
	"log"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	log.Printf("Receive message body from client: %s", in.Body)

}
