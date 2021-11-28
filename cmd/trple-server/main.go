package main

import (
	"context"
	"fmt"
	"log"
	"net"

	productV1 "github.com/brettmostert/trple-proto-go/trple/product/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Println("server is meant to do something")

	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	productV1.RegisterProductServiceServer(server, &productServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type productServiceServer struct {
	productV1.UnimplementedProductServiceServer
}

func (s *productServiceServer) GetProductById(ctx context.Context, req *productV1.GetProductByIdRequest) (*productV1.GetProductByIdResponse, error) {
	return &productV1.GetProductByIdResponse{}, nil
}
