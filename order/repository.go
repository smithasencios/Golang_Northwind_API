package order

import (
	"context"
	"database/sql"
)

type Repository interface {
	InsertOrder(ctx context.Context, params *addOrderRequest) (int64, error)
	InsertOrderDetail(ctx context.Context, params *addOrderDetailRequest) (int64, error)
	GetOrders(ctx context.Context, params *getOrdersRequest) ([]*OrderListItem, error)
	// GetOrderDetail(ctx context.Context, orderId *int64) ([]*OrderDetailListItem, error)
	GetTotalOrders() (int64, error)
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
	result, err := repo.db.Exec(sql, params.OrderID, params.ProductID, params.Quantity, params.UnitPrice)

	if err != nil {
		panic(err.Error())
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) GetOrders(ctx context.Context, params *getOrdersRequest) ([]*OrderListItem, error) {
	const sql = `SELECT o.id,o.customer_id,o.order_date,o.status_id,os.status_name,
	CONCAT(c.first_name,' ',c.last_name) as customer_name
	FROM orders o
	INNER JOIN orders_status os ON o.status_id = os.id
	INNER JOIN customers c ON o.customer_id = c.id
	LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)

	if err != nil {
		panic(err.Error())
	}
	var orders []*OrderListItem

	for results.Next() {
		order := &OrderListItem{}
		err = results.Scan(&order.ID, &order.CustomerID, &order.OrderDate, &order.StatusId, &order.StatusName, &order.Customer)
		if err != nil {
			panic(err.Error())
		}
		orderDetail, err := GetOrderDetail(repo, &order.ID)
		if err != nil {
			panic(err.Error())
		}

		order.Data = orderDetail
		orders = append(orders, order)
	}
	return orders, nil
}
func (repo *repository) GetTotalOrders() (int64, error) {
	const sql = "SELECT COUNT(*) FROM orders"
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err.Error())
	}
	return total, nil
}
func GetOrderDetail(repo *repository, orderId *int64) ([]*OrderDetailListItem, error) {
	const sql = `SELECT order_id,quantity,unit_price,p.product_name
	FROM order_details od
	INNER JOIN products p ON od.product_id = p.id
	WHERE od.order_id = ?`
	results, err := repo.db.Query(sql, orderId)

	if err != nil {
		panic(err.Error())
	}
	var orders []*OrderDetailListItem

	for results.Next() {
		order := &OrderDetailListItem{}
		err = results.Scan(&order.OrderId, &order.Quantity, &order.UnitPrice, &order.ProductName)
		if err != nil {
			panic(err.Error())
		}
		orders = append(orders, order)
	}
	return orders, nil
}
