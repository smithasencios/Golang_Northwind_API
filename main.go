package main

import (
	"log"
	"net/http"

	"github.com/Golang_Northwind_API/auth"
	"github.com/Golang_Northwind_API/customer"
	"github.com/Golang_Northwind_API/employee"
	"github.com/Golang_Northwind_API/order"
	"github.com/Golang_Northwind_API/product"
	"github.com/Golang_Northwind_API/shared"
	"github.com/codegangsta/negroni"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbInstance := shared.InitDB()

	var (
		employeeRepo = employee.NewRepository(dbInstance)
		productRepo  = product.NewRepository(dbInstance)
		customerRepo = customer.NewRepository(dbInstance)
		orderRepo    = order.NewRepository(dbInstance)
	)
	var (
		employeeService employee.Service
		productService  product.Service
		customerService customer.Service
		orderService    order.Service
		authService     auth.Service
	)
	employeeService = employee.New(employeeRepo)
	productService = product.New(productRepo)
	customerService = customer.New(customerRepo)
	orderService = order.New(orderRepo)
	authService = auth.New()

	r := chi.NewRouter()
	r.Use(shared.GetCors().Handler)

	jwtMiddleware := auth.GetJwtMiddleware()

	r.Mount("/orders", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(order.MakeHTTPHandler(orderService)),
	))
	r.Mount("/products", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(product.MakeHTTPHandler(productService)),
	))
	r.Mount("/employees", employee.MakeHTTPHandler(employeeService))
	r.Mount("/customers", customer.MakeHTTPHandler(customerService))
	r.Mount("/auth", auth.MakeHTTPHandler(authService))

	server := http.ListenAndServe(":3000", r)
	log.Fatal(server)
}
