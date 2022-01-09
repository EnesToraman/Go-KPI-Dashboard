package model

type Employee struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Title   string `json:"title"`
	Salary  int    `json:"salary"`
}

type EmployeeDTO struct {
	Title             string `json:"title"`
	NumberOfEmployees int    `json:"numberOfEmployees"`
}

type EmployeeSalary struct {
	Title     string `json:"title"`
	AvgSalary int    `json:"salary"`
}
