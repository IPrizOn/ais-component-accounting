package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateComponent(conn *pgx.Conn, id int, comp string, desc string, price int) error {
	query := `UPDATE component 
				SET type = $2, description = $3, price = $4
					WHERE id = $1;`

	_, err := conn.Exec(context.Background(), query, id, comp, desc, price)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCustomer(conn *pgx.Conn, id int, name string, phone string, email string, address string) error {
	query := `UPDATE customer 
				SET name = $2, phone = $3, email = $4, address = $5
					WHERE id = $1;`

	_, err := conn.Exec(context.Background(), query, id, name, phone, email, address)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSale(conn *pgx.Conn, id int, compID int, custID int, count int) error {
	query := `UPDATE sale 
				SET component_id = $2, customer_id = $3, count = $4 
					WHERE id = $1;`

	_, err := conn.Exec(context.Background(), query, id, compID, custID, count)
	if err != nil {
		return err
	}

	return nil
}
