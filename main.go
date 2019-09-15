package main

import (
	"log"
	"net/http"

	"github.com/Golang_Northwind_API/auth"
	"github.com/Golang_Northwind_API/customer"
	_ "github.com/Golang_Northwind_API/docs"
	"github.com/Golang_Northwind_API/employee"
	"github.com/Golang_Northwind_API/order"
	"github.com/Golang_Northwind_API/product"
	"github.com/Golang_Northwind_API/shared"
	"github.com/codegangsta/negroni"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Northwind API
// @version 1.0
// @description Northwind API

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email code4humans@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.Print("Logging in Go!")
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

	// use ginSwagger middleware to serve the API docs
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("../swagger/doc.json"),
	))

	server := http.ListenAndServe(":3000", r)
	log.Fatal(server)
}
