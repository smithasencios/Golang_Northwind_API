package customer

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetCustomers(ctx context.Context, params *getCustomersRequest) ([]*Customer, error)
	GetTotalCustomers() (int64, error)
}
type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
func (repo *repository) GetCustomers(ctx context.Context, params *getCustomersRequest) ([]*Customer, error) {
	const sql = `SELECT c.id,c.first_name,c.last_name,address,business_phone,city,company
				FROM customers c
				LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)

	if err != nil {
		panic(err.Error())
	}
	var customers []*Customer

	for results.Next() {
		customer := &Customer{}
		err = results.Scan(&customer.ID, &customer.First_Name, &customer.Last_Name, &customer.Address, &customer.Business_Phone,
			&customer.City, &customer.Company)
		if err != nil {
			panic(err.Error())
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
func (repo *repository) GetTotalCustomers() (int64, error) {
	const sql = "SELECT COUNT(*) FROM customers"
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err.Error())
	}
	return total, nil
}
