package main

import (
	"ecomm/internal/controller"
	"ecomm/proto"
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}

	conn, err := grpc.NewClient("localhost:8081", options...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewApiServiceClient(conn)
	productHandler := controller.NewHandler(client)

	router := controller.NewRouter(productHandler)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
