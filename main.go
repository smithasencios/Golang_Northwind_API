package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/u/Go_Angular/backend/customer"
	"github.com/u/Go_Angular/backend/employee"
	"github.com/u/Go_Angular/backend/product"
)

func main() {
	//lambda.Start(handler)
	dbInstance := initDB()
	var (
		employeeRepo = employee.NewRepository(dbInstance)
		productRepo  = product.NewRepository(dbInstance)
		customerRepo = customer.NewRepository(dbInstance)
	)
	var (
		employeeService employee.Service
		productService  product.Service
		customerService customer.Service
	)
	employeeService = employee.New(employeeRepo)
	productService = product.New(productRepo)
	customerService = customer.New(customerRepo)

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

	r.Mount("/employees", employee.MakeHTTPHandler(employeeService))
	r.Mount("/products", product.MakeHTTPHandler(productService))
	r.Mount("/customers", customer.MakeHTTPHandler(customerService))

	server := http.ListenAndServe(":8080", r)
	log.Fatal(server)
}

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "demouser:cofebe@tcp(127.0.0.1:3306)/northwind")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	return db
}
func handler() (string, error) {
	return "Welcome to serverless world", nil
}
