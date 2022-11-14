-- Create Operation
INSERT INTO dbo.Employee (UserName, LastName, FirstName, Phone, City)
VALUES ('megaraj.selvam', 'Selvam', 'Megaraj', '123456789', 'Mumbai');

-- Retrieve Operation
select * from dbo.Employee

-- Update Operation
UPDATE dbo.Employee
SET FirstName = 'Mega', LastName = 'S', Phone = '1234567890', City = 'Mumbai'
WHERE EmpID = 1 AND UserName = 'megaraj.selvam';

-- Delete Operation
DELETE FROM Employee WHERE EmpID=1 AND UserName = 'megaraj.selvam';