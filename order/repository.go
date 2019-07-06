package order

import (
	"context"
	"database/sql"
)
type Repository interface {
	InsertOrder(ctx context.Context, params *addOrderRequest) (int64, error)
	InsertOrderDetail(ctx context.Context, params *addOrderDetailRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
func (repo *repository) InsertOrder(ctx context.Context, params *addOrderRequest) (int64, error) {
	const sql = `
	INSERT INTO orders
	(customer_id ,order_date)
	VALUES(?,?)`
	result, err := repo.db.Exec(sql, params.CustomerID, params.OrderDate)

	if err != nil {
		panic(err.Error())
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) InsertOrderDetail(ctx context.Context, params *addOrderDetailRequest) (int64, error) {
	const sql = `
	INSERT INTO order_details
	(order_id ,product_id,quantity,unit_price)
	VALUES(?,?,?,?)`
	result, err := repo.db.Exec(sql, params.OrderID, params.ProductID,params.Quantity,params.UnitPrice)

	if err != nil {
		panic(err.Error())
	}

	id, _ := result.LastInsertId()
	return id, nil
}