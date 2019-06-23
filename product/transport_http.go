package product

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductsHandler := kithttp.NewServer(
		makeGetProductsEndpoint(s),
		getProductsRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	addProductHandler := kithttp.NewServer(
		makeAddProductEndpoint(s),
		addProductRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addProductHandler)

	deleteProductHandler := kithttp.NewServer(
		makeDeleteProjectEndpoint(s),
		getDeleteProjectRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)

	getProductByIDHandler := kithttp.NewServer(
		makeGetProductByIDEndpoint(s),
		getProjectByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/{id}", getProductByIDHandler)

	updateProductHandler := kithttp.NewServer(
		makeUpdateProductEndpoint(s),
		updateProductRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/", updateProductHandler)

	getBestSellersHandler := kithttp.NewServer(
		makeBestSellersEndpoint(s),
		getBestSellerRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/bestSellers", getBestSellersHandler)

	return r
}

func getProductsRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
func addProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	fmt.Printf("%+v\n", request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
func getDeleteProjectRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return deleteProjectRequest{
		ProjectID: chi.URLParam(r, "id"),
	}, nil
}
func getProjectByIDRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getProductByIDRequest{
		ProductID: chi.URLParam(r, "id"),
	}, nil
}

func updateProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := updateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	fmt.Printf("%+v\n", request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
func getBestSellerRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getBestSellersRequest{}, nil
}
