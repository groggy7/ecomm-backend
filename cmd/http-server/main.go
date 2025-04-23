package main

import (
	"context"
	"ecomm/internal/controller"
	"ecomm/internal/controller/router"
	"ecomm/internal/repository"
	"ecomm/internal/usecases"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	connString := os.Getenv("CONN_STRING")
	if connString == "" {
		log.Fatal("CONN_STRING is not set")
	}

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}

	productRepo := repository.NewProductRepo(pool)
	productUsecase := usecases.NewProductUsecase(productRepo)
	productHandler := controller.NewProductHandler(productUsecase)

	router := router.NewRouter(productHandler)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
