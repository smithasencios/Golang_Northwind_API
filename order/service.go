package order

import (
	"context"
)

type Service interface {
	InsertOrder(ctx context.Context, params *addOrderRequest) (int64, error)
	GetOrders(ctx context.Context, params *getOrdersRequest) (*OrderList, error)
}

type service struct {
	repo Repository
}

func New(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) InsertOrder(ctx context.Context, params *addOrderRequest) (int64, error) {
	orderId, err := s.repo.InsertOrder(ctx, params)
	if err != nil {
		panic(err.Error())
	}
	for _, detail := range params.OrderDetails {
		detail.OrderID = orderId
		_, err := s.repo.InsertOrderDetail(ctx, &detail)
		if err != nil {
			panic(err.Error())
		}
	}
	return orderId, err
}
func (s *service) GetOrders(ctx context.Context, params *getOrdersRequest) (*OrderList, error) {
	orders, err := s.repo.GetOrders(ctx, params)
	if err != nil {
		panic(err.Error())
	}
	totalOrders, err := s.repo.GetTotalOrders()
	if err != nil {
		panic(err.Error())
	}
	return &OrderList{Data: orders, TotalRecords: totalOrders}, err
}
