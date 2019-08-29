package auth

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getPermissionsByUserId struct {
	UserID string
}

func makeGetPermissionsEndpoint(s Service) endpoint.Endpoint {
	getPermissionsEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getPermissionsByUserId)
		r, err := s.GetPermissions(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getPermissionsEndpoint
}
