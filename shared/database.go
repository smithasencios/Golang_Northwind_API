package shared

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

func InitDB() *sql.DB {
	conexionString := join(os.Getenv("NORTHWIND_DB_USER"), ":", os.Getenv("NORTHWIND_DB_PASSWORD"), "@tcp(", os.Getenv("DATABASE_HOST"), ")/", os.Getenv("NORTHWIND_DB_DATABASE"))
	fmt.Println(conexionString)
	log.Print(conexionString)
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
