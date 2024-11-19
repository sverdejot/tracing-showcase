package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	v1 "github.com/sverdejot/tracing-showcase/pkg/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	addr = ":8088"
)

type service struct {
	v1.UnimplementedForwarderServiceServer
}

func (s *service) Forward(ctx context.Context, msg *v1.ForwardRequest) (*v1.ForwardResponse, error) {
	log.Println("got: ", msg.Message)
	return nil, nil
}

func main() {
	if p := os.Getenv("GRPC_PORT"); p != "" {
		addr = fmt.Sprintf(":%s", p)
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Unable to instantiate listener: %v", err)
	}
	defer lis.Close()

	srv := grpc.NewServer()
	service := &service{}

	reflection.Register(srv)
	v1.RegisterForwarderServiceServer(srv, service)

	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
