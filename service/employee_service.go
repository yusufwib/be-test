package service

import (
	"context"

	"github.com/yusufwib/be-test/models/demployee"
	"github.com/yusufwib/be-test/repository"
)

type IEmployeeService interface {
	FindAll(ctx context.Context) ([]demployee.EmployeeResponse, error)
	FindByID(ctx context.Context, ID uint64) (demployee.EmployeeResponse, error)
	Create(ctx context.Context, req demployee.EmployeeRequest) error
	Update(ctx context.Context, req demployee.EmployeeRequest) error
	Delete(ctx context.Context, ID uint64) error
}

type EmployeeService struct {
	EmployeeRepository repository.EmployeeRepository
}

func NewEmployeeService(i repository.EmployeeRepository) EmployeeService {
	return EmployeeService{
		EmployeeRepository: i,
	}
}

func (i *EmployeeService) FindAll(ctx context.Context) ([]demployee.EmployeeResponse, error) {
	return i.EmployeeRepository.FindAll(ctx)
}

func (i *EmployeeService) FindByID(ctx context.Context, ID uint64) (demployee.EmployeeResponse, error) {
	return i.EmployeeRepository.FindByID(ctx, ID)
}

func (i *EmployeeService) Create(ctx context.Context, req demployee.EmployeeRequest) error {
	return i.EmployeeRepository.Create(ctx, req.ToEmployeeData())
}

func (i *EmployeeService) Update(ctx context.Context, req demployee.EmployeeRequest) error {
	return i.EmployeeRepository.Update(ctx, req.ToEmployeeData())
}

func (i *EmployeeService) Delete(ctx context.Context, ID uint64) error {
	return i.EmployeeRepository.Delete(ctx, ID)
}
