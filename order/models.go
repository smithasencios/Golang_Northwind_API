package order

type (
	OrderList struct {
		Data         []*OrderListItem `json:"data"`
		TotalRecords int64            `json:"totalRecords"`
	}

	OrderListItem struct {
		ID         int64                  `json:"order_id"`
		CustomerID int                    `json:"customer_id"`
		OrderDate  string                 `json:"order_date"`
		StatusId   string                 `json:"status_id"`
		StatusName string                 `json:"status_name"`
		Customer   string                 `json:"customer"`
		Data       []*OrderDetailListItem `json:"data"`
	}

	OrderDetailList struct {
		Data []*OrderDetailListItem `json:"data"`
	}
	OrderDetailListItem struct {
		OrderId     int     `json:"order_id"`
		Quantity    float64 `json:"quantity"`
		UnitPrice   float64 `json:"unit_price"`
		ProductName string  `json:"product_name"`
	}
)
