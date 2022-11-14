package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	DRIVERNAME = "mssql"
	SQLSERVER  = "localhost"
	PORT       = 1433
	USERNAME   = "sa"
	PASSWORD   = "reallyStrongPwd123"
	DATABASE   = "employeeDB"
)

func getConnectionString() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", SQLSERVER, USERNAME, PASSWORD, PORT, DATABASE)
}

func GetSQLConnection() (db *sql.DB) {
	conn, err := sql.Open(DRIVERNAME, getConnectionString())
	if err != nil {
		log.Fatal(err.Error())
	}
	return conn
}
