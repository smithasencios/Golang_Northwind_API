package order

import (
	"context"
)

type Service interface {
	InsertOrder(ctx context.Context, params *addOrderRequest) (int64, error)
	UpdateOrder(ctx context.Context, params *addOrderRequest) (int64, error)
	GetOrders(ctx context.Context, params *getOrdersRequest) (*OrderList, error)
	GetOrderById(ctx context.Context, params *getOrderByIdRequest) (*OrderListItem, error)
	DeleteOrderDetail(ctx context.Context, params *deleteOrderDetailRequest) (int64, error)
	DeleteOrder(ctx context.Context, params *deleteOrderRequest) (int64, error)
}

type service struct {
	repo Repository
}

func New(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
func (s *service) UpdateOrder(ctx context.Context, params *addOrderRequest) (int64, error) {
	orderId, err := s.repo.UpdateOrder(ctx, params)
	if err != nil {
		panic(err.Error())
	}
	for _, detail := range params.OrderDetails {
		detail.OrderID = orderId
		if detail.ID == 0 {
			_, err = s.repo.InsertOrderDetail(ctx, &detail)
		} else {
			_, err = s.repo.UpdateOrderDetail(ctx, &detail)
		}

		if err != nil {
			panic(err.Error())
		}
	}
	return orderId, err
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
	totalOrders, err := s.repo.GetTotalOrders(ctx, params)
	if err != nil {
		panic(err.Error())
	}
	return &OrderList{Data: orders, TotalRecords: totalOrders}, err
}

func (s *service) GetOrderById(ctx context.Context, params *getOrderByIdRequest) (*OrderListItem, error) {
	order, err := s.repo.GetOrderById(ctx, params)
	if err != nil {
		panic(err.Error())
	}
	return order, err
}

func (s *service) DeleteOrderDetail(ctx context.Context, params *deleteOrderDetailRequest) (int64, error) {
	return s.repo.DeleteOrderDetail(ctx, params)
}

func (s *service) DeleteOrder(ctx context.Context, params *deleteOrderRequest) (int64, error) {

	_, err := s.repo.DeleteOrderDetailByIdOrderId(ctx, params)
	if err != nil {
		panic(err.Error())
	}
	return s.repo.DeleteOrder(ctx, params)
}
