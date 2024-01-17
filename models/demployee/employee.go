package demployee

import (
	"time"
)

type Employee struct {
	ID        uint64    `json:"id" gorm:"column:id;type:serial primary_key"`
	FirstName string    `json:"first_name" gorm:"column:first_name;type:text"`
	LastName  string    `json:"last_name" gorm:"column:last_name;type:text"`
	Email     string    `json:"email" gorm:"column:email;type:text"`
	HireDate  time.Time `json:"hire_date" gorm:"column:hire_date;type:date"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Employee) TableName() string {
	return "employees.employees"
}
