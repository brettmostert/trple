package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	productV1 "github.com/brettmostert/trple-proto-product-go/trple/product/v1"

	"google.golang.org/genproto/googleapis/type/money"
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
	p := &productV1.GetProductByIdResponse{}

	p.Product = &productV1.Product{}
	p.Product.Id = "123" + strconv.Itoa(int(req.GetId()))
	p.Product.Name = "123"
	p.Product.Howmuch = &money.Money{}
	p.Product.Howmuch.CurrencyCode = "ZAR"
	p.Product.Howmuch.Units = 100
	p.Product.Howmuch.Nanos = 50

	return p, nil
}
