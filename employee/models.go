package employee

type (
	Employee struct {
		ID            int    `json:"id"`
		LastName      string `json:"last_name"`
		FirstName     string `json:"first_name"`
		Company       string `json:"company"`
		EmailAddress  string `json:"email_address"`
		JobTitle      string `json:"job_title"`
		BusinessPhone string `json:"business_phone"`
		HomePhone     string `json:"home_phone"`
		MobilePhone   string `json:"mobile_phone"`
		FaxNumber     string `json:"fax_number"`
		Address       string `json:"address"`
	}
	EmployeeList struct {
		Data         []*Employee `json:"data"`
		TotalRecords int64       `json:"totalRecords"`
	}
	BestEmployee struct {
		ID          int    `json:"id"`
		TotalVentas int    `json:"totalVentas"`
		LastName    string `json:"last_name"`
		FirstName   string `json:"first_name"`
	}
)
