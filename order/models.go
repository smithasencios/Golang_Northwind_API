package order

type (
	OrderList struct {
		Data         []*OrderListItem `json:"data"`
		TotalRecords int64            `json:"totalRecords"`
	}

	OrderListItem struct {
		ID         int    `json:"order_id"`
		CustomerID int    `json:"customer_id"`
		OrderDate  string `json:"order_date"`
		StatusId   string `json:"status_id"`
		StatusName string `json:"status_name"`
		Customer   string `json:"customer"`
	}
)
