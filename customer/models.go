package customer

type (
	Customer struct {
		ID             int    `json:"id"`
		First_Name     string `json:"first_name"`
		Last_Name      string `json:"last_name"`
		Address        string `json:"address"`
		Business_Phone string `json:"business_phone"`
		City           string `json:"city"`
		Company        string `json:"company"`
	}
	CustomerList struct {
		Data         []*Customer `json:"data"`
		TotalRecords int64       `json:"totalRecords"`
	}
)
