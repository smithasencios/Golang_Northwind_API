package order

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
	addOrderHandler := kithttp.NewServer(
		makeAddProductEndpoint(s),
		addOrderRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addOrderHandler)
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
