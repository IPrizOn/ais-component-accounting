package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateComponent(conn *pgx.Conn) error {
	query := `UPDATE component 
				SET type = $2, description = $3, price = $4
					WHERE id = $1;`

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCustomer(conn *pgx.Conn) error {
	query := `UPDATE customer 
				SET name = $2, phone = $3, email = $4, address = $5
					WHERE id = $1;`

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSale(conn *pgx.Conn) error {
	query := `UPDATE sale 
				SET component_id = $2, customer_id = $3, count = $4 
					WHERE id = $1;`

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}
