package employee

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetEmployees(ctx context.Context, params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
	GetEmployeeByID(ctx context.Context, params *getEmployeeByIDRequest) (*Employee, error)
	UpdateEmployee(ctx context.Context, params *updateEmployeeRequest) (int64, error)
	DeleteEmployee(ctx context.Context, params *deleteEmployeeRequest) (int64, error)
	GetBestEmployee(ctx context.Context, params *getBestEmployeeRequest) (*BestEmployee, error)
	InsertEmployee(ctx context.Context, params *addEmployeeRequest) (int64, error)
}
type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
func (repo *repository) GetEmployees(ctx context.Context, params *getEmployeesRequest) ([]*Employee, error) {
	const sql = `SELECT id,first_name,last_name,company,email_address,job_title,business_phone,home_phone,
				COALESCE(mobile_phone,''),fax_number,address
				FROM employees
				LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)

	if err != nil {
		panic(err.Error())
	}
	var employees []*Employee

	for results.Next() {
		employee := &Employee{}
		err = results.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company,
			&employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone,
			&employee.MobilePhone, &employee.FaxNumber, &employee.Address)
		if err != nil {
			panic(err.Error())
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (repo *repository) GetTotalEmployees() (int64, error) {
	const sql = "SELECT COUNT(*) FROM employees"
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err.Error())
	}
	return total, nil
}

func (repo *repository) GetEmployeeByID(ctx context.Context, params *getEmployeeByIDRequest) (*Employee, error) {
	const sql = `SELECT id,first_name,last_name,company,email_address,job_title,business_phone,home_phone,
				COALESCE(mobile_phone,''),fax_number,address
				FROM employees
				WHERE id=?`

	row := repo.db.QueryRow(sql, params.EmployeeID)

	employee := &Employee{}

	err := row.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company,
		&employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone,
		&employee.MobilePhone, &employee.FaxNumber, &employee.Address)

	if err != nil {
		panic(err.Error())
	}
	return employee, nil
}

func (repo *repository) UpdateEmployee(ctx context.Context, params *updateEmployeeRequest) (int64, error) {
	const sql = `
			UPDATE employees
			SET first_name = ?,
			last_name = ?,
			company=?,
			address=?,
			business_phone=?,
			email_address=?,
			fax_number=?,
			home_phone=?,
			job_title=?,
			mobile_phone=?
			WHERE id = ?`
	result, err := repo.db.Exec(sql, params.First_Name, params.Last_Name, params.Company,
		params.Address, params.Business_Phone, params.Email_Address, params.Fax_Number, params.Home_Phone,
		params.Job_Title, params.Mobile_Phone, params.ID)

	if err != nil {
		panic(err.Error())
	}

	id, _ := result.LastInsertId()
	return id, nil
}
func (repo *repository) InsertEmployee(ctx context.Context, params *addEmployeeRequest) (int64, error) {
	const sql = `
	INSERT INTO employees
	(first_name ,last_name,company,address,business_phone,email_address,
	fax_number,home_phone,job_title,mobile_phone)
	VALUES(?,?,?,?,?,?,?,?,?,?)`
	result, err := repo.db.Exec(sql, params.First_Name, params.Last_Name, params.Company,
		params.Address, params.Business_Phone, params.Email_Address, params.Fax_Number, params.Home_Phone,
		params.Job_Title, params.Mobile_Phone)

	if err != nil {
		panic(err.Error())
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) DeleteEmployee(ctx context.Context, params *deleteEmployeeRequest) (int64, error) {
	const sql = `DELETE FROM employees WHERE id = ?`
	result, err := repo.db.Exec(sql, params.EmployeeID)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return count, nil
}

func (repo *repository) GetBestEmployee(ctx context.Context, params *getBestEmployeeRequest) (*BestEmployee, error) {
	const sql = `SELECT e.id,count(e.id) as totalVentas,e.first_name,e.last_name
				FROM orders o
				INNER JOIN employees e  ON o.employee_id = e.id
				GROUP BY o.employee_id
				ORDER BY totalVentas desc
				limit 1`
	row := repo.db.QueryRow(sql)

	employee := &BestEmployee{}

	err := row.Scan(&employee.ID, &employee.TotalVentas, &employee.FirstName, &employee.LastName)

	if err != nil {
		panic(err.Error())
	}
	return employee, nil
}
