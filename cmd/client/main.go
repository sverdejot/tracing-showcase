package main

import (
	"context"
	"fmt"
	"log"
	"os"

	v1 "github.com/sverdejot/tracing-showcase/pkg/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	target := os.Getenv("TARGET")
	port := os.Getenv("PORT")
	if target == "" || port == "" {
		log.Fatal("Required variables 'TARGET' or 'PORT' not set")
	}

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", target, port),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatalf("Unable to instantiate client: %v", err)
	}
	client := v1.NewForwarderServiceClient(conn)

	_, err = client.Forward(context.Background(), &v1.ForwardRequest{Message: "Hello world!"}, grpc.WaitForReady(true))
	if err != nil {
		log.Fatal(err)
	}
}
