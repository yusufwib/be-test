package demployee

import "time"

type EmployeeRequest struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	HireDate  string `json:"hire_date" validate:"required"`
}

func (e EmployeeRequest) ToEmployeeData() Employee {
	hireDate, _ := time.Parse("2006-01-02", e.HireDate)
	return Employee{
		ID:        e.ID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Email:     e.Email,
		HireDate:  hireDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
