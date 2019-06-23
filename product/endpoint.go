package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductsRequest struct {
	Limit  int
	Offset int
}
type getAddProductRequest struct {
	Category      string
	Description   string
	List_Price    string
	Standard_Cost string
	Product_Code  string
	Product_Name  string
}
type updateProductRequest struct {
	ID            int
	Category      string
	Description   string
	List_Price    float32
	Standard_Cost float32
	Product_Code  string
	Product_Name  string
}
type deleteProjectRequest struct {
	ProjectID string
}
type getProductByIDRequest struct {
	ProductID string
}
type getBestSellersRequest struct{}

func makeGetProductsEndpoint(s Service) endpoint.Endpoint {
	getProductsEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		r, err := s.GetProducts(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getProductsEndpoint
}
func makeAddProductEndpoint(s Service) endpoint.Endpoint {
	addProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		r, err := s.InsertProduct(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return addProductEndpoint
}
func makeDeleteProjectEndpoint(s Service) endpoint.Endpoint {
	deleteProjectEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProjectRequest)
		r, err := s.DeleteProject(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return deleteProjectEndpoint
}
func makeGetProductByIDEndpoint(s Service) endpoint.Endpoint {
	getProductByIDEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		r, err := s.GetProductByID(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getProductByIDEndpoint
}
func makeUpdateProductEndpoint(s Service) endpoint.Endpoint {
	updateProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		r, err := s.UpdateProduct(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return updateProductEndpoint
}
func makeBestSellersEndpoint(s Service) endpoint.Endpoint {
	getBestSellersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getBestSellersRequest)
		r, err := s.GetBestSellers(ctx, &req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return getBestSellersEndpoint
}
