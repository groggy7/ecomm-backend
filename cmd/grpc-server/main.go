package main

import (
	"context"
	"ecomm/internal/repository"
	"ecomm/internal/service"
	"ecomm/proto"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	connString := os.Getenv("CONN_STRING")
	if connString == "" {
		log.Fatal("CONN_STRING is not set")
	}

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}

	productRepo := repository.NewRepository(pool)
	productService := service.NewService(productRepo)

	server := grpc.NewServer()
	proto.RegisterApiServiceServer(server, productService)

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server started on port 8081")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
