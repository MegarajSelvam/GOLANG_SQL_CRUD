package handler

import (
	"GO_SQL_CRUD/contract"
	"GO_SQL_CRUD/helper"
	"GO_SQL_CRUD/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var employee contract.IEmpCRUDOperator

func EmployeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	operation := strings.ToLower(r.Method)

	switch operation {
	case "post":
		createEmployeeHandler(w, r)
	case "get":
		retrieveEmployeeHandler(w, r)
	case "put":
		updateEmployeeHandler(w, r)
	case "delete":
		deleteEmployeeHandler(w, r)
	default:
		message := fmt.Sprintf("%s operation is not supported", r.Method)
		w.WriteHeader(405)
		w.Write(helper.ErrorToBytes(types.CreateErrorMessage(message)))
	}
}

func createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	err := setEmployee(r)
	if err != nil {
		writeResult(w, helper.AnyToErrorBytes(err), 400)
		return
	}
	result, status := employee.Create()
	if !status {
		writeResult(w, helper.AnyToErrorBytes(result), 400)
		return
	}
	result, _ = employee.RetrieveByUserName()
	writeResult(w, helper.AnyToEmployeeBytes(result), 201)
}

func retrieveEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		getAllEmployeesHandler(w)
	} else {
		getEmployeesByIdHandler(w, id)
	}
}

func getAllEmployeesHandler(w http.ResponseWriter) {
	result, status := employee.Retrieve()
	if !status {
		writeResult(w, helper.AnyToErrorBytes(result), 400)
		return
	}
	writeResult(w, helper.AnyToEmployeeListBytes(result), 200)
}

func getEmployeesByIdHandler(w http.ResponseWriter, id string) {
	empId, err := helper.StringToInt(id)
	if err != nil {
		writeResult(w, helper.AnyToErrorBytes(err), 400)
		return
	}
	emp := employee.(types.Employee)
	emp.EmpId = empId
	result, status := emp.RetrieveByEmpId()
	if !status {
		writeResult(w, helper.AnyToErrorBytes(result), 400)
		return
	}
	writeResult(w, helper.AnyToEmployeeBytes(result), 200)
}

func updateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	err := setEmployee(r)
	if err != nil {
		writeResult(w, helper.AnyToErrorBytes(err), 400)
		return
	}
	result, status := employee.Update()
	if !status {
		writeResult(w, helper.AnyToErrorBytes(result), 400)
		return
	}
	result, _ = employee.RetrieveByUserName()
	writeResult(w, helper.AnyToEmployeeBytes(result), 202)
}

func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	empId, err := helper.StringToInt(id)
	if err != nil {
		writeResult(w, helper.AnyToErrorBytes(err), 400)
		return
	}
	emp := employee.(types.Employee)
	emp.EmpId = empId
	result, status := emp.Delete()
	if !status {
		writeResult(w, helper.AnyToErrorBytes(result), 400)
		return
	}
	writeResult(w, nil, 204)
}

func setEmployee(r *http.Request) any {
	requestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return types.CreateErrorMessage(err.Error())
	}
	emp := employee.(types.Employee)
	err = json.Unmarshal(requestBytes, &emp)
	if err != nil {
		return types.CreateErrorMessage(err.Error())
	}
	employee = emp
	return nil
}

func writeResult(w http.ResponseWriter, response []byte, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write(response)
}

func init() {
	employee = types.Employee{}
}
