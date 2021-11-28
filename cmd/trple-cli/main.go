package main

import (
	"context"
	"fmt"
	"log"

	productV1 "github.com/brettmostert/trple-proto-go/trple/product/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to PetStoreService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)

	productClient := productV1.NewProductServiceClient(conn)

	response, err := productClient.GetProductById(context.Background(), &productV1.GetProductByIdRequest{
		Id: 7})
	if err != nil {
		return fmt.Errorf("failed to PutPet: %w", err)
	}
	response.GetResult()
	log.Println("Successfully PutPet", response)
	return nil
}
