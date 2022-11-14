package main

import (
	"GO_SQL_CRUD/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/employee", handler.EmployeeHandler)
	log.Fatal(http.ListenAndServe(":7002", nil))
}
