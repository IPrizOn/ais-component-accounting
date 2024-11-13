package main

import (
	"ais/client"
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
)

var (
	conn *pgx.Conn
	once sync.Once
)

func main() {
	connection := GetInstanceConnect()

	defer func() {
		err := connection.Close(context.Background())
		if err != nil {
			log.Panic(err)
		}
	}()

	client.Start(connection)
}

func GetInstanceConnect() *pgx.Conn {
	connectionString := "postgresql://postgres:13799731@localhost:5432/accounting_components?sslmode=disable"

	once.Do(func() {
		var err error

		conn, err = pgx.Connect(context.Background(), connectionString)
		if err != nil {
			log.Panic(err)
		}
	})

	return conn
}
