package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yusufwib/be-test/models/demployee"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	return "employees"
}

func TestEmployeeRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	err = db.AutoMigrate(&Employee{})
	if err != nil {
		t.Fatalf("Error running migrations: %v", err)
	}

	repo := NewEmployeeRepository(db, "employees")

	testEmployee := demployee.Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		HireDate:  time.Now(),
	}

	t.Run("Create and FindByID", func(t *testing.T) {
		err := repo.Create(context.Background(), testEmployee)
		if err != nil {
			t.Fatalf("failed to create employee: %v", err)
		}

		foundEmployee, err := repo.FindByID(context.Background(), testEmployee.ID)
		if err != nil {
			t.Fatalf("failed to find employee by ID: %v", err)
		}

		if foundEmployee.FirstName != testEmployee.FirstName || foundEmployee.LastName != testEmployee.LastName {
			t.Fatalf("expected %+v, got %+v", testEmployee, foundEmployee)
		}
	})

	t.Run("FindAll", func(t *testing.T) {
		employees, err := repo.FindAll(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, employees)
	})

	t.Run("FindByID", func(t *testing.T) {
		foundEmployee, err := repo.FindByID(context.Background(), testEmployee.ID)
		assert.NoError(t, err)
		assert.Equal(t, testEmployee.FirstName, foundEmployee.FirstName)
		assert.Equal(t, testEmployee.LastName, foundEmployee.LastName)
	})

	t.Run("Update", func(t *testing.T) {
		testEmployee.FirstName = "UpdatedJohn"
		err = repo.Update(context.Background(), testEmployee)
		assert.NoError(t, err)

		updatedEmployee, err := repo.FindByID(context.Background(), testEmployee.ID)
		assert.NoError(t, err)
		assert.Equal(t, "UpdatedJohn", updatedEmployee.FirstName)
	})

	t.Run("Delete", func(t *testing.T) {
		err = repo.Delete(context.Background(), testEmployee.ID)
		assert.NoError(t, err)
	})
}
