package product

import (
	"context"
)

type Service interface {
	GetProducts(ctx context.Context, request *getProductsRequest) (*ProductList, error)
	InsertProduct(ctx context.Context, params *getAddProductRequest) (int64, error)
	DeleteProject(ctx context.Context, params *deleteProjectRequest) (int64, error)
	GetProductByID(ctx context.Context, params *getProductByIDRequest) (*Product, error)
	UpdateProduct(ctx context.Context, params *updateProductRequest) (int64, error)
	GetBestSellers(ctx context.Context, params *getBestSellersRequest) (*ProductTopResponse, error)
}

type service struct {
	repo Repository
}

func New(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetProducts(ctx context.Context, request *getProductsRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(ctx, request)

	if err != nil {
		panic(err.Error())
	}
	totalProducts, err := s.repo.GetTotalProducts()
	if err != nil {
		panic(err.Error())
	}
	return &ProductList{Data: products, TotalRecords: totalProducts}, err

}
func (s *service) InsertProduct(ctx context.Context, params *getAddProductRequest) (int64, error) {
	return s.repo.InsertProduct(ctx, params)
}
func (s *service) DeleteProject(ctx context.Context, params *deleteProjectRequest) (int64, error) {
	return s.repo.DeleteProject(ctx, params)
}
func (s *service) GetProductByID(ctx context.Context, params *getProductByIDRequest) (*Product, error) {
	return s.repo.GetProductByID(ctx, params)
}
func (s *service) UpdateProduct(ctx context.Context, params *updateProductRequest) (int64, error) {
	return s.repo.UpdateProduct(ctx, params)
}
func (s *service) GetBestSellers(ctx context.Context, request *getBestSellersRequest) (*ProductTopResponse, error) {
	products, err := s.repo.GetBestSellers(ctx, request)

	if err != nil {
		panic(err.Error())
	}
	totalVentas, err := s.repo.GetTotalVentas()
	if err != nil {
		panic(err.Error())
	}
	return &ProductTopResponse{Data: products, TotalVentas: totalVentas}, err
}
