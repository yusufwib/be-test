package demployee

import "time"

type EmployeeResponse struct {
	ID        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	HireDate  time.Time `json:"hire_date"`
}

func (e EmployeeResponse) IsEmpty() bool {
	return e == EmployeeResponse{}
}
