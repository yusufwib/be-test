package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/yusufwib/be-test/models/demployee"
	"gorm.io/gorm"
)

type IEmployeeRepository interface {
	FindAll(ctx context.Context) ([]demployee.EmployeeResponse, error)
	FindByID(ctx context.Context, ID uint64) (demployee.EmployeeResponse, error)
	Create(ctx context.Context, req demployee.Employee) error
	Update(ctx context.Context, req demployee.Employee) error
	Delete(ctx context.Context, ID uint64) error
}

type EmployeeRepository struct {
	DB        *gorm.DB
	TableName string
}

func NewEmployeeRepository(DB *gorm.DB, tableName string) EmployeeRepository {
	if tableName == "" {
		tableName = EmployeeTable
	}
	return EmployeeRepository{DB, tableName}
}

const EmployeeTable = "employees.employees"

func (repo *EmployeeRepository) session(ctx context.Context) *gorm.DB {
	trx, ok := ctx.Value("pg").(*gorm.DB)
	if !ok {
		return repo.DB
	}
	return trx
}

func (repo *EmployeeRepository) FindAll(ctx context.Context) (res []demployee.EmployeeResponse, err error) {
	trx := repo.session(ctx)
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err = trx.WithContext(ctxWT).Table(repo.TableName).Find(&res).Error; err != nil {
		return res, fmt.Errorf("err while get employees: %w", err)
	}
	return
}

func (repo *EmployeeRepository) FindByID(ctx context.Context, ID uint64) (res demployee.EmployeeResponse, err error) {
	trx := repo.session(ctx)
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err = trx.WithContext(ctxWT).Table(repo.TableName).Where("id = ?", ID).Find(&res).Error; err != nil {
		return res, fmt.Errorf("err while get employee by id: %w", err)
	}
	return
}

func (repo *EmployeeRepository) Create(ctx context.Context, req demployee.Employee) (err error) {
	trx := repo.session(ctx)
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err = trx.WithContext(ctxWT).Table(repo.TableName).Create(&req).Error; err != nil {
		return fmt.Errorf("err while create employee: %w", err)
	}
	return
}

func (repo *EmployeeRepository) Update(ctx context.Context, req demployee.Employee) (err error) {
	trx := repo.session(ctx)
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err = trx.WithContext(ctxWT).Table(repo.TableName).Where("id = ?", req.ID).Updates(map[string]interface{}{
		"first_name": req.FirstName,
		"last_name":  req.LastName,
		"email":      req.Email,
		"hire_date":  req.HireDate,
		"updated_at": req.UpdatedAt,
	}).Error; err != nil {
		return fmt.Errorf("err while update employee: %w", err)
	}
	return
}

func (repo *EmployeeRepository) Delete(ctx context.Context, ID uint64) (err error) {
	trx := repo.session(ctx)
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err = trx.WithContext(ctxWT).Table(repo.TableName).Where("id = ?", ID).Delete(demployee.Employee{}).Error; err != nil {
		return fmt.Errorf("err while delete employee: %w", err)
	}
	return nil
}
