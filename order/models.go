package order

type (
	Order struct {
		OrderHeader  OrderHeader    `json:"orderHeader"`
		OrderDetails []OrderDetails `json:"orderDetails"`
	}
	OrderHeader struct {
		ID        int    `json:"order_id"`
		OrderDate string `json:"order_date"`
	}
	OrderDetails struct {
		ID        int     `json:"id"`
		ProductID int     `json:"product_id"`
		Quantity  int     `json:"quantity"`
		UnitPrice float64 `json:"unit_price"`
	}
)
