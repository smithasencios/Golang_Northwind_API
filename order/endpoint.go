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
