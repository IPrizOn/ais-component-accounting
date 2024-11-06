package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateComponent(conn *pgx.Conn) error {
	query := `INSERT INTO component (type, description, price)
				VALUES ($1, $2, $3);`

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func CreateCustomer(conn *pgx.Conn) error {
	query := `INSERT INTO customer (name, phone, email, address)
				VALUES ($1, $2, $3, $4);`

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func CreateSale(conn *pgx.Conn) error {
	query := `INSERT INTO sale (component_id, customer_id, count)
				VALUES ($1, $2, $3);`

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}
