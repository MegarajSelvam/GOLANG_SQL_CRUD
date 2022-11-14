package types

import (
	"GO_SQL_CRUD/config"
	"database/sql"
	"fmt"
)

type Employee struct {
	EmpId     int    `json:"id"`
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	City      string `json:"city"`
}

const (
	CREATE_QUERY               = "INSERT INTO dbo.Employee (UserName, LastName, FirstName, Phone, City) VALUES (?, ?, ?, ?, ?)"
	RETRIEVE_QUERY             = "SELECT * FROM dbo.Employee ORDER BY EmpID"
	RETRIEVE_BY_USERNAME_QUERY = "SELECT * FROM dbo.Employee WHERE UserName = ?"
	RETRIEVE_BY_ID_QUERY       = "SELECT * FROM dbo.Employee WHERE EmpID = ?"
	UPDATE_QUERY               = "UPDATE dbo.Employee SET FirstName = ?, LastName = ?, Phone = ?, City = ? WHERE EmpID = ?;"
	DELETE_QUERY               = "DELETE FROM Employee WHERE EmpID = ?"
)

func (e Employee) Create() (any, bool) {
	conn := config.GetSQLConnection()
	defer conn.Close()
	createEmployee, err := conn.Prepare(CREATE_QUERY)
	if err != nil {
		return CreateErrorMessage(err.Error()), false
	}
	_, err = createEmployee.Exec(e.UserName, e.LastName, e.FirstName, e.Phone, e.City)
	if err != nil {
		return CreateErrorMessage(err.Error()), false
	}
	return e, true
}

func (e Employee) Retrieve() (any, bool) {
	conn := config.GetSQLConnection()
	defer conn.Close()
	employeeRows, err := conn.Query(RETRIEVE_QUERY)
	if err != nil {
		return CreateErrorMessage(err.Error()), false
	}
	result, status := scanEmployee(employeeRows, 0)
	return result, status
}

func (e Employee) RetrieveByUserName() (any, bool) {
	conn := config.GetSQLConnection()
	defer conn.Close()
	employeeRows, err := conn.Query(RETRIEVE_BY_USERNAME_QUERY, e.UserName)
	if err != nil {
		return CreateErrorMessage(err.Error()), false
	}
	result, status := scanEmployee(employeeRows, 1)
	if result == nil && !status {
		return CreateErrorMessage(fmt.Sprintf("No employee found with User Name %s", e.UserName)), false
	}
	return result, status
}

func (e Employee) RetrieveByEmpId() (any, bool) {
	conn := config.GetSQLConnection()
	defer conn.Close()
	employeeRows, err := conn.Query(RETRIEVE_BY_ID_QUERY, e.EmpId)
	if err != nil {
		return CreateErrorMessage(err.Error()), false
	}
	result, status := scanEmployee(employeeRows, 1)
	if result == nil && !status {
		return CreateErrorMessage(fmt.Sprintf("No employee found with Id %d", e.EmpId)), false
	}
	return result, status
}

func (e Employee) Update() (any, bool) {
	conn := config.GetSQLConnection()
	defer conn.Close()
	updateEmployee, err := conn.Prepare(UPDATE_QUERY)
	if err != nil {
		return CreateErrorMessage(err.Error()), false
	}
	result, err := updateEmployee.Exec(e.FirstName, e.LastName, e.Phone, e.City, e.EmpId)
	if err != nil {
		return CreateErrorMessage(err.Error()), false
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return CreateErrorMessage(err.Error()), false
	} else if rows <= 0 {
		return CreateErrorMessage(fmt.Sprintf("No record found with employee id: %d", e.EmpId)), false
	}
	return e, true
}

func (e Employee) Delete() (any, bool) {
	conn := config.GetSQLConnection()
	defer conn.Close()
	deleteEmployee, err := conn.Prepare(DELETE_QUERY)
	if err != nil {
		return CreateErrorMessage(err.Error()), false
	}
	result, err := deleteEmployee.Exec(e.EmpId)
	if err != nil {
		return CreateErrorMessage(err.Error()), false
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return CreateErrorMessage(err.Error()), false
	} else if rows <= 0 {
		return CreateErrorMessage(fmt.Sprintf("No record found for employee id: %d", e.EmpId)), false
	}
	return e, true
}

// limit 1 means return one employee, zero means return all fetched rows
func scanEmployee(rows *sql.Rows, limit int) (any, bool) {
	empArray := []Employee{}
	emp := Employee{}
	for rows.Next() {
		var UserName, LastName, FirstName, Phone, City string
		var EmpId int
		err := rows.Scan(&EmpId, &UserName, &LastName, &FirstName, &Phone, &City)
		if err != nil {
			return err, false
		}
		emp.EmpId = EmpId
		emp.UserName = UserName
		emp.FirstName = FirstName
		emp.LastName = LastName
		emp.Phone = Phone
		emp.City = City
		empArray = append(empArray, emp)
	}
	if limit == 1 {
		if emp.EmpId == 0 {
			return nil, false
		}
		return emp, true
	}
	return empArray, true
}
