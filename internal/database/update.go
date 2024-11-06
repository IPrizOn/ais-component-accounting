package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateComponent(conn *pgx.Conn) error {
	query := ``

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCustomer(conn *pgx.Conn) error {
	query := ``

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSale(conn *pgx.Conn) error {
	query := ``

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}
