package order

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()
	addOrderHandler := kithttp.NewServer(
		makeAddProductEndpoint(s),
		addOrderRequestDecoder,
		kithttp.EncodeJSONResponse,
	)

	r.Method(http.MethodPost, "/", addOrderHandler)

	getOrdersHandler := kithttp.NewServer(
		makeGetOrdersEndpoint(s),
		getOrdersRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/paginated", getOrdersHandler)

	getOrderByIdHandler := kithttp.NewServer(
		makeGetOrderByIdEndpoint(s),
		getOrderByIdRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/{id}", getOrderByIdHandler)

	return r
}

func addOrderRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	fmt.Printf("%+v\n", request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
func getOrdersRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getOrdersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
func getOrderByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	orderId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	return getOrderByIdRequest{
		orderId: orderId,
	}, nil
}
