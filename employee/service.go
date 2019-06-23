package employee

import (
	"context"
)

type Service interface {
	GetEmployees(ctx context.Context, request *getEmployeesRequest) (*EmployeeList, error)
	GetEmployeesByID(ctx context.Context, request *getEmployeeByIDRequest) (*Employee, error)
	UpdateEmployee(ctx context.Context, params *updateEmployeeRequest) (int64, error)
	DeleteEmployee(ctx context.Context, params *deleteEmployeeRequest) (int64, error)
	GetBestEmployee(ctx context.Context, params *getBestEmployeeRequest) (*BestEmployee, error)
	InsertEmployee(ctx context.Context, params *addEmployeeRequest) (int64, error)
}

type service struct {
	repo Repository
}

func New(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetEmployees(ctx context.Context, request *getEmployeesRequest) (*EmployeeList, error) {
	employees, err := s.repo.GetEmployees(ctx, request)

	if err != nil {
		panic(err.Error())
	}
	totalEmployees, err := s.repo.GetTotalEmployees()
	if err != nil {
		panic(err.Error())
	}
	return &EmployeeList{Data: employees, TotalRecords: totalEmployees}, err

}
func (s *service) GetEmployeesByID(ctx context.Context, request *getEmployeeByIDRequest) (*Employee, error) {
	return s.repo.GetEmployeeByID(ctx, request)
}
func (s *service) UpdateEmployee(ctx context.Context, params *updateEmployeeRequest) (int64, error) {
	return s.repo.UpdateEmployee(ctx, params)
}
func (s *service) DeleteEmployee(ctx context.Context, params *deleteEmployeeRequest) (int64, error) {
	return s.repo.DeleteEmployee(ctx, params)
}
func (s *service) GetBestEmployee(ctx context.Context, params *getBestEmployeeRequest) (*BestEmployee, error) {
	return s.repo.GetBestEmployee(ctx, params)
}
func (s *service) InsertEmployee(ctx context.Context, params *addEmployeeRequest) (int64, error) {
	return s.repo.InsertEmployee(ctx, params)
}
