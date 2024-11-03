package main

import (
	"ais/internal/client"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := connectDB()
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		err := conn.Close(context.Background())
		if err != nil {
			log.Panic(err)
		}
	}()

	client.Start(conn)
}

func connectDB() (*pgx.Conn, error) {
	connectionString := "postgresql://postgres:13799731@localhost:5432/accounting_components?sslmode=disable"

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Panic(err)
	}

	return conn, nil
}
