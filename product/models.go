package product

type (
	Product struct {
		ID           int     `json:"id"`
		ProductCode  string  `json:"product_code"`
		ProductName  string  `json:"product_name"`
		Description  string  `json:"description"`
		StandardCost float64 `json:"standard_cost"`
		ListPrice    float64 `json:"list_price"`
		ReorderLevel int     `json:"reorder_level"`
		TargetLevel  int     `json:"target_level"`
		Category     string  `json:"category"`
	}
	ProductList struct {
		Data         []*Product `json:"data"`
		TotalRecords int64      `json:"totalRecords"`
	}
	ProductTop struct {
		ID          int     `json:"id"`
		ProductName string  `json:"product_name"`
		Vendidos    float64 `json:"vendidos"`
	}
	ProductTopResponse struct {
		Data        []*ProductTop `json:"data"`
		TotalVentas float64       `json:"totalVentas"`
	}
)
