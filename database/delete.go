package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteComponent(conn *pgx.Conn, id int) error {
	query := `DELETE
				FROM component
					WHERE id = $1;`

	_, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCustomer(conn *pgx.Conn, id int) error {
	query := `DELETE
				FROM customer
					WHERE id = $1;`

	_, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSale(conn *pgx.Conn, id int) error {
	query := `DELETE
				FROM sale
					WHERE id = $1;`

	_, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}
