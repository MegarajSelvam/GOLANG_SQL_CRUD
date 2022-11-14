CREATE TABLE Employee (
    EmpID int NOT NULL IDENTITY(1, 1),
    UserName VARCHAR(255) NOT NULL,
    LastName varchar(255),
    FirstName varchar(255) NOT NULL,
    Phone varchar(255),
    City varchar(255),
    UNIQUE (UserName),
    PRIMARY KEY (EmpID)
);