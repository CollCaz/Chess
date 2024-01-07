package main

import (
	"Chess/internal/app"
	data "Chess/internal/models"
	"Chess/internal/server"
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("could not connect to db")
	}
	server := server.NewServer()

	app := &app.App
	app.Models = data.NewModels(pool)

	err = server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
