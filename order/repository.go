package order

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	UpdateOrder(ctx context.Context, params *addOrderRequest) (int64, error)
	UpdateOrderDetail(ctx context.Context, params *addOrderDetailRequest) (int64, error)
	InsertOrder(ctx context.Context, params *addOrderRequest) (int64, error)
	InsertOrderDetail(ctx context.Context, params *addOrderDetailRequest) (int64, error)
	GetOrders(ctx context.Context, params *getOrdersRequest) ([]*OrderListItem, error)
	GetOrderById(ctx context.Context, params *getOrderByIdRequest) (*OrderListItem, error)
	GetTotalOrders(ctx context.Context, params *getOrdersRequest) (int64, error)
	DeleteOrderDetail(ctx context.Context, params *deleteOrderDetailRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) UpdateOrder(ctx context.Context, params *addOrderRequest) (int64, error) {
	const sql = `
	UPDATE orders
	SET customer_id = ?
	WHERE id = ? `
	_, err := repo.db.Exec(sql, params.CustomerID, params.ID)

	if err != nil {
		panic(err.Error())
	}
	return params.ID, nil
}

func (repo *repository) UpdateOrderDetail(ctx context.Context, params *addOrderDetailRequest) (int64, error) {
	const sql = `
	UPDATE order_details
	SET quantity = ?,
		unit_price = ?
	WHERE id = ?`
	_, err := repo.db.Exec(sql, params.Quantity, params.UnitPrice, params.ID)

	if err != nil {
		panic(err.Error())
	}
	return params.ID, nil
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
	var filter string

	if params.Status != nil {
		filter += fmt.Sprintf(" AND o.status_id  = %v ", params.Status.(float64))
	}
	if params.Date_From != nil && params.Date_To == nil {
		filter += fmt.Sprintf(" AND o.order_date  >= '%v' ", params.Date_From.(string))
	}
	if params.Date_From == nil && params.Date_To != nil {
		filter += fmt.Sprintf(" AND o.order_date  <= '%v' ", params.Date_To.(string))
	}
	if params.Date_From != nil && params.Date_To != nil {
		filter += fmt.Sprintf(" AND o.order_date  between '%v' and '%v' ", params.Date_From.(string), params.Date_To.(string))
	}

	var sql = `SELECT o.id,
	o.customer_id,
	o.order_date,
	o.status_id,
	os.status_name,
	CONCAT(c.first_name,' ',c.last_name) as customer_name	
	FROM orders o
	INNER JOIN orders_status os ON o.status_id = os.id
	INNER JOIN customers c ON o.customer_id = c.id
	WHERE 1=1 `

	sql = sql + filter + "LIMIT ? OFFSET ?"

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
	return orders, err
}

func (repo *repository) GetTotalOrders(ctx context.Context, params *getOrdersRequest) (int64, error) {
	var total int64
	var filter string

	if params.Status != nil {
		filter += fmt.Sprintf(" AND o.status_id  = %v ", params.Status.(float64))
	}
	if params.Date_From != nil && params.Date_To == nil {
		filter += fmt.Sprintf(" AND o.order_date  >= '%v' ", params.Date_From.(string))
	}
	if params.Date_From == nil && params.Date_To != nil {
		filter += fmt.Sprintf(" AND o.order_date  <= '%v' ", params.Date_To.(string))
	}
	if params.Date_From != nil && params.Date_To != nil {
		filter += fmt.Sprintf(" AND o.order_date  between '%v' and '%v' ", params.Date_From.(string), params.Date_To.(string))
	}

	var sql = "SELECT COUNT(*) FROM orders o WHERE 1=1 " + filter

	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err.Error())
	}
	return total, nil
}

func (repo *repository) GetOrderById(ctx context.Context, params *getOrderByIdRequest) (*OrderListItem, error) {
	var sql = `SELECT o.id,o.customer_id,o.order_date,o.status_id,os.status_name,
	CONCAT(c.first_name,' ',c.last_name) as customer_name,
	c.company,
	c.address,
	c.business_phone,
	c.city
	FROM orders o
	INNER JOIN orders_status os ON o.status_id = os.id
	INNER JOIN customers c ON o.customer_id = c.id
	WHERE o.id = ? `

	order := &OrderListItem{}

	row := repo.db.QueryRow(sql, params.orderId)
	err := row.Scan(&order.ID, &order.CustomerID, &order.OrderDate, &order.StatusId, &order.StatusName, &order.Customer,
		&order.Company, &order.Address, &order.Phone, &order.City)
	if err != nil {
		panic(err.Error())
	}
	orderDetail, err := GetOrderDetail(repo, &order.ID)
	if err != nil {
		panic(err.Error())
	}

	order.Data = orderDetail

	return order, err
}

func GetOrderDetail(repo *repository, orderId *int64) ([]*OrderDetailListItem, error) {
	const sql = `SELECT order_id,od.id,quantity,unit_price,p.product_name,product_id
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
		err = results.Scan(&order.OrderId, &order.ID, &order.Quantity, &order.UnitPrice, &order.ProductName, &order.ProductId)
		if err != nil {
			panic(err.Error())
		}
		orders = append(orders, order)
	}
	return orders, nil
}
func (repo *repository) DeleteOrderDetail(ctx context.Context, params *deleteOrderDetailRequest) (int64, error) {
	const sql = `DELETE FROM order_details WHERE id = ?`
	result, err := repo.db.Exec(sql, params.OrderDetailID)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return count, nil
}
