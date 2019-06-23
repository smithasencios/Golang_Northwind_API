package employee

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getEmployeesRequest struct {
	Limit  int
	Offset int
}
type getBestEmployeeRequest struct{}

type getEmployeeByIDRequest struct {
	EmployeeID string
}
type deleteEmployeeRequest struct {
	EmployeeID string
}
type updateEmployeeRequest struct {
	ID             int
	Address        string
	Business_Phone string
	Company        string
	Email_Address  string
	Fax_Number     string
	First_Name     string
	Home_Phone     string
	Job_Title      string
	Last_Name      string
	Mobile_Phone   string
}
type addEmployeeRequest struct {
	Address        string
	Business_Phone string
	Company        string
	Email_Address  string
	Fax_Number     string
	First_Name     string
	Home_Phone     string
	Job_Title      string
	Last_Name      string
	Mobile_Phone   string
}

func makeGetEmployeesEndpoint(s Service) endpoint.Endpoint {
	getEmployeesEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)
		r, err := s.GetEmployees(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getEmployeesEndpoint
}
func makeGetEmployeeByIDEndpoint(s Service) endpoint.Endpoint {
	getEmployeeByIDEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeByIDRequest)
		r, err := s.GetEmployeesByID(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getEmployeeByIDEndpoint
}
func makeUpdateEmployeeEndpoint(s Service) endpoint.Endpoint {
	getUpdateEmployeeEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEmployeeRequest)
		r, err := s.UpdateEmployee(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getUpdateEmployeeEndpoint
}
func makeDeleteEmployeeEndpoint(s Service) endpoint.Endpoint {
	deleteEmployeeEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteEmployeeRequest)
		r, err := s.DeleteEmployee(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return deleteEmployeeEndpoint
}
func makeGetBestEmployeeEndpoint(s Service) endpoint.Endpoint {
	getBestEmployeeEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getBestEmployeeRequest)
		r, err := s.GetBestEmployee(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getBestEmployeeEndpoint
}
func makeInsertEmployeeEndpoint(s Service) endpoint.Endpoint {
	getInsertEmployeeEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addEmployeeRequest)
		r, err := s.InsertEmployee(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getInsertEmployeeEndpoint
}
