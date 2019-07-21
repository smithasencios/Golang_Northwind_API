package main

import (
	"database/sql"
	"log"
	"net/http"

	"fmt"
	"os"
	"strings"

	"github.com/Golang_Northwind_API/auth"
	"github.com/Golang_Northwind_API/customer"
	"github.com/Golang_Northwind_API/employee"
	"github.com/Golang_Northwind_API/order"
	"github.com/Golang_Northwind_API/product"
	"github.com/codegangsta/negroni"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbInstance := initDB()
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
	)
	employeeService = employee.New(employeeRepo)
	productService = product.New(productRepo)
	customerService = customer.New(customerRepo)
	orderService = order.New(orderRepo)

	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	jwtMiddleware := auth.GetJwtMiddleware()

	r.Mount("/orders", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(order.MakeHTTPHandler(orderService)),
	))
	r.Mount("/employees", employee.MakeHTTPHandler(employeeService))
	r.Mount("/products", product.MakeHTTPHandler(productService))
	r.Mount("/customers", customer.MakeHTTPHandler(customerService))
	// r.Mount("/orders", order.MakeHTTPHandler(orderService))

	server := http.ListenAndServe(":3000", r)
	log.Fatal(server)
}

func initDB() *sql.DB {
	conexionString := join(os.Getenv("NORTHWIND_DB_USER"), ":", os.Getenv("NORTHWIND_DB_PASSWORD"), "@tcp(", os.Getenv("DATABASE_HOST"), ")/", os.Getenv("NORTHWIND_DB_DATABASE"))
	fmt.Println(conexionString)
	db, err := sql.Open("mysql", conexionString)
	//root:lfda@tcp(localhost:3306)/northwind

	if err != nil {

		panic(err.Error())
	}
	// defer db.Close()
	return db
}
func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
