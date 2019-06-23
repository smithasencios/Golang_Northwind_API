package employee

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

	getEmployeesHandler := kithttp.NewServer(
		makeGetEmployeesEndpoint(s),
		getEmployeesRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/paginated", getEmployeesHandler)

	getEmployeeByIDHandler := kithttp.NewServer(
		makeGetEmployeeByIDEndpoint(s),
		getEmployeeByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/{id}", getEmployeeByIDHandler)

	updateEmployeesHandler := kithttp.NewServer(
		makeUpdateEmployeeEndpoint(s),
		getUpdateEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/", updateEmployeesHandler)

	deleteEmployeesHandler := kithttp.NewServer(
		makeDeleteEmployeeEndpoint(s),
		getDeleteEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodDelete, "/{id}", deleteEmployeesHandler)

	getBestEmployeeHandler := kithttp.NewServer(
		makeGetBestEmployeeEndpoint(s),
		getBestEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/best", getBestEmployeeHandler)

	addEmployeesHandler := kithttp.NewServer(
		makeInsertEmployeeEndpoint(s),
		getAddEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addEmployeesHandler)

	return r
}

func getEmployeesRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
func getEmployeeByIDRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getEmployeeByIDRequest{
		EmployeeID: chi.URLParam(r, "id"),
	}, nil
}
func getUpdateEmployeeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := updateEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
func getDeleteEmployeeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return deleteEmployeeRequest{
		EmployeeID: chi.URLParam(r, "id"),
	}, nil
}
func getBestEmployeeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getBestEmployeeRequest{}, nil
}
func getAddEmployeeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := addEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	fmt.Printf("%+v\n", request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
