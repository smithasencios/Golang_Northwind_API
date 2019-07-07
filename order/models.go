package order

type (
	OrderList struct {
		Data         []*OrderHeader `json:"data"`
		TotalRecords int64          `json:"totalRecords"`
	}
	Order struct {
		OrderHeader  OrderHeader    `json:"orderHeader"`
		OrderDetails []OrderDetails `json:"orderDetails"`
	}
	OrderHeader struct {
		ID           int            `json:"order_id"`
		CustomerID   int            `json:"customer_id"`
		OrderDate    string         `json:"order_date"`
		OrderDetails []OrderDetails `json:"order_details"`
	}
	OrderDetails struct {
		ID        int     `json:"id"`
		OrderID   int     `json:"order_id"`
		ProductID int     `json:"product_id"`
		Quantity  int     `json:"quantity"`
		UnitPrice float64 `json:"unit_price"`
	}
)
