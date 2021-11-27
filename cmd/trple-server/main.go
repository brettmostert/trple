package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	productV1 "github.com/brettmostert/trple-proto-go/product/v1"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("server is meant to do something")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
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
}

type productServiceServer struct {
	productV1.UnimplementedProductServiceServer
}

func (s *productServiceServer) PutPet(ctx context.Context, req *productV1.FetchResponseRequest) (*productV1.FetchResponseResponse, error) {
	name := req.GetName()
	petType := req.GetPetType()
	log.Println("Got a request to create a", petType, "named", name)

	return &productV1.FetchResponseResponse{}, nil
}
