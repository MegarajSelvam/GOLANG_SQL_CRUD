# #######################################################################################
#                                 MICROSOFT SQL SERVER CRUD OPERATION USING GOLANG
# #######################################################################################

# # TOOLS REQUIRED:

1) Docker Desktop
2) Visual Studio Code
3) Golang
4) Azure data Studio
# #######################################################################################

# INSTALLING SQL SERVER INSTANCES THROUGH DOCKER IMAGES

For Windows/For Mac (Intel Chip):
```bash
docker run -d --name sql -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=reallyStrongPwd123' -p 1433:1433 mcr.microsoft.com/mssql/server:2019-latest
```

For Mac (M1 Chip): 
```bash
docker run -e "ACCEPT_EULA=1" -e "MSSQL_SA_PASSWORD=reallyStrongPwd123" -e "MSSQL_PID=Developer" -e "MSSQL_USER=SA" -p 1433:1433 -d --name=sql mcr.microsoft.com/azure-sql-edge
```
# #######################################################################################

# CONNECT SQL SERVER INSTANCES USING AZURE DATA STUDIO

Server Name:
    For Windows:(local) 
    For Mac: localhost 
UserName: sa
Password: reallyStrongPwd123

# #######################################################################################

# EXECUTE THE DDL COMMANDS IN FOLLOWING ORDER FROM ssms_query FOLDER

1) CreateEmployeeDB.sql
2) CreateEmployeeTable.sql

# #######################################################################################

# # TO RUN THE SOLUTION

To install all dependencies
```bash
go get github.com/denisenkom/go-mssqldb
```

To run the solution
```bash
go run .
```
# #######################################################################################

# # TO TEST THE API

Use postman collection attached in postman folder
# #######################################################################################
