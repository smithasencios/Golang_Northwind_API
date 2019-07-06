package order

import (
	"context"
)

type Service interface {
	InsertOrder(ctx context.Context, params *addOrderRequest) (int64, error)
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
