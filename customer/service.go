package customer

import (
	"context"
)

type Service interface {
	GetCustomers(ctx context.Context, request *getCustomersRequest) (*CustomerList, error)
}

type service struct {
	repo Repository
}

func New(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetCustomers(ctx context.Context, request *getCustomersRequest) (*CustomerList, error) {
	customers, err := s.repo.GetCustomers(ctx, request)

	if err != nil {
		panic(err.Error())
	}
	totalCustomers, err := s.repo.GetTotalCustomers()
	if err != nil {
		panic(err.Error())
	}
	return &CustomerList{Data: customers, TotalRecords: totalCustomers}, err

}
