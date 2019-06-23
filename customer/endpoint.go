package customer

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getCustomersRequest struct {
	Limit  int
	Offset int
}

func makeGetcustomersEndpoint(s Service) endpoint.Endpoint {
	getCustomersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCustomersRequest)
		r, err := s.GetCustomers(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getCustomersEndpoint
}
