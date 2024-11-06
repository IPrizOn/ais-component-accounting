package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateComponent(conn *pgx.Conn) error {
	query := ``

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func CreateCustomer(conn *pgx.Conn) error {
	query := ``

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func CreateSale(conn *pgx.Conn) error {
	query := ``

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}
