package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Arlodee69/expense-tracker/internal/expense"
	"github.com/Arlodee69/expense-tracker/internal/httpapi"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	/* $env:DATABASE_URL = "postgres://user:password@localhost:5432/expense_tracker?sslmode=disable" */
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("Specify the environment variable \"DATABASE_URL\"")
	}
	log.Println("Environment variable obtained")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Panic(err)
	}
	defer pool.Close()
	err = pool.Ping(ctx)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Подключение к PostgreSQL проверено")

	storage := expense.NewPostgresStore(pool)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /expenses", func(w http.ResponseWriter, r *http.Request) {
		httpapi.HandlerExpenses(w, r, storage)
	})
	mux.HandleFunc("POST /expenses", func(w http.ResponseWriter, r *http.Request) {
		httpapi.HandlerExpenses(w, r, storage)
	})
	log.Println("server started: http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
