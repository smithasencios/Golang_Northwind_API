package order

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addOrderRequest struct {
	ID           int64
	CustomerID   int
	OrderDate    string
	OrderDetails []addOrderDetailRequest
}
type addOrderDetailRequest struct {
	ID        int64
	OrderID   int64
	ProductID int64
	Quantity  int64
	UnitPrice float64
}
type getOrdersRequest struct {
	Limit     int
	Offset    int
	Status    interface{}
	Date_From interface{}
	Date_To   interface{}
}
type getOrderByIdRequest struct {
	orderId int64
}

type deleteOrderDetailRequest struct {
	OrderDetailID string
}

type deleteOrderRequest struct {
	OrderID string
}

func makeAddProductEndpoint(s Service) endpoint.Endpoint {
	addOrderEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		r, err := s.InsertOrder(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return addOrderEndpoint
}
func makeUpdateOrderEndpoint(s Service) endpoint.Endpoint {
	updateOrderEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		r, err := s.UpdateOrder(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return updateOrderEndpoint
}
func makeGetOrdersEndpoint(s Service) endpoint.Endpoint {
	getOrdersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersRequest)
		r, err := s.GetOrders(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getOrdersEndpoint
}

func makeGetOrderByIdEndpoint(s Service) endpoint.Endpoint {
	getOrderByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderByIdRequest)
		r, err := s.GetOrderById(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getOrderByIdEndpoint
}

func makeDeleteOrderDetailEndpoint(s Service) endpoint.Endpoint {
	deleteOrderDeleteEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderDetailRequest)
		r, err := s.DeleteOrderDetail(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return deleteOrderDeleteEndpoint
}

func makeDeleteOrderEndpoint(s Service) endpoint.Endpoint {
	deleteOrderDeleteEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderRequest)
		r, err := s.DeleteOrder(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return deleteOrderDeleteEndpoint
}
