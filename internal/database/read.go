package database

import (
	"ais/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

func GetComponentsList(conn *pgx.Conn) ([]domain.Component, error) {
	query := `SELECT * 
				FROM component`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		componentsList []domain.Component
		component      domain.Component
	)

	for rows.Next() {
		err = rows.Scan(&component.ID, &component.Type, &component.Description, &component.Price)
		if err != nil {
			return nil, err
		}

		componentsList = append(componentsList, component)
	}

	return componentsList, nil
}

func GetCustomersList(conn *pgx.Conn) ([]domain.Customer, error) {
	query := `SELECT *
				FROM customer`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		customerList []domain.Customer
		customer     domain.Customer
	)

	for rows.Next() {
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.Email, &customer.Address)
		if err != nil {
			return nil, err
		}

		customerList = append(customerList, customer)
	}

	return customerList, nil
}

func GetSalesList(conn *pgx.Conn) ([]domain.Sale, error) {
	query := `SELECT *
				FROM sale`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		salesList []domain.Sale
		sale      domain.Sale
	)

	for rows.Next() {
		err = rows.Scan(&sale.ID, &sale.ComponentID, &sale.CustomerID, &sale.Count)
		if err != nil {
			return nil, err
		}

		salesList = append(salesList, sale)
	}

	return salesList, nil
}

func GetPersonsList(conn *pgx.Conn) ([]domain.Person, error) {

	return nil, nil
}

func IsUserExist(conn *pgx.Conn) error {

	return nil
}
