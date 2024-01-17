package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/yusufwib/be-test/models/demployee"
	"github.com/yusufwib/be-test/service"
	"github.com/yusufwib/be-test/utils/mvalidator"
	"github.com/yusufwib/be-test/utils/trace_id"

	mlog "github.com/yusufwib/be-test/utils/logger"

	"github.com/labstack/echo/v4"
)

type (
	IEmployeeHander interface {
		FindAll(ctx echo.Context) error
		FindByID(ctx echo.Context) error
		Create(ctx echo.Context) error
		Update(ctx echo.Context) error
		Delete(ctx echo.Context) error
	}

	EmployeeHandler struct {
		Context         context.Context
		Logger          mlog.Logger
		Validator       mvalidator.Validator
		EmployeeService service.EmployeeService
	}
)

func NewEmployeeHandler(
	context context.Context,
	logger mlog.Logger,
	validator mvalidator.Validator,
	employeeService service.EmployeeService,
) IEmployeeHander {
	return &EmployeeHandler{
		Context:         context,
		Logger:          logger,
		Validator:       validator,
		EmployeeService: employeeService,
	}
}

// FindAll swagger
// @Summary Get all employees
// @Tags employees
// @Produce json
// @Success 200 {object} []demployee.Employee
// @Router /employees [get]
func (i *EmployeeHandler) FindAll(ctx echo.Context) error {
	traceID := trace_id.GetID(ctx)
	usecaseContext := trace_id.SetIDx(ctx.Request().Context(), traceID)

	i.Logger.InfoT(traceID, "get all employees")

	if employees, err := i.EmployeeService.FindAll(usecaseContext); err != nil {
		return ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	} else if len(employees) == 0 {
		return ErrorResponse(ctx, http.StatusNotFound, "No employees found", nil)
	} else {
		return SuccessResponse(ctx, http.StatusOK, employees)
	}
}

// FindByID swagger
// @Summary Get an employee by ID
// @Tags employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} demployee.Employee
// @Router /employees/{id} [get]
func (i *EmployeeHandler) FindByID(ctx echo.Context) error {
	traceID := trace_id.GetID(ctx)
	usecaseContext := trace_id.SetIDx(ctx.Request().Context(), traceID)

	ID, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
	if err != nil {
		return ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	i.Logger.InfoT(traceID, "get employee by id", mlog.Any("id", ID))

	if employee, err := i.EmployeeService.FindByID(usecaseContext, ID); err != nil {
		return ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	} else if employee.IsEmpty() {
		return ErrorResponse(ctx, http.StatusNotFound, "No employees found", nil)
	} else {
		return SuccessResponse(ctx, http.StatusOK, employee)
	}
}

// Create swagger
// @Summary Create a new employee
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body demployee.EmployeeRequest true "Employee data"
// @Success 201 {object} demployee.Employee
// @Router /employees [post]
func (i *EmployeeHandler) Create(ctx echo.Context) error {
	traceID := trace_id.GetID(ctx)
	usecaseContext := trace_id.SetIDx(ctx.Request().Context(), traceID)

	var req demployee.EmployeeRequest
	if err := ctx.Bind(&req); err != nil {
		return ErrorResponse(ctx, http.StatusBadRequest, "bad request", nil)
	}

	if mapErr, err := i.Validator.Struct(req); err != nil {
		return ErrorResponse(ctx, http.StatusBadRequest, "invalid payload", mapErr)
	}

	i.Logger.InfoT(traceID, "create employee", mlog.Any("payload", req))

	if err := i.EmployeeService.Create(usecaseContext, req); err != nil {
		return ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	} else {
		return SuccessResponse(ctx, http.StatusCreated, nil)
	}
}

// Update swagger
// @Summary Update an employee by ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body demployee.EmployeeRequest true "Employee data"
// @Success 200 {object} demployee.Employee
// @Router /employees/{id} [put]
func (i *EmployeeHandler) Update(ctx echo.Context) error {
	traceID := trace_id.GetID(ctx)
	usecaseContext := trace_id.SetIDx(ctx.Request().Context(), traceID)

	var req demployee.EmployeeRequest
	if err := ctx.Bind(&req); err != nil {
		return ErrorResponse(ctx, http.StatusBadRequest, "bad request", nil)
	}

	if mapErr, err := i.Validator.Struct(req); err != nil {
		return ErrorResponse(ctx, http.StatusBadRequest, "invalid payload", mapErr)
	}

	ID, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
	if err != nil {
		return ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	i.Logger.InfoT(traceID, "update employee by id", mlog.Any("id", ID), mlog.Any("payload", req))

	req.ID = ID
	if err := i.EmployeeService.Update(usecaseContext, req); err != nil {
		return ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	} else {
		return SuccessResponse(ctx, http.StatusOK, nil)
	}
}

// Delete swagger
// @Summary Delete an employee by ID
// @Tags employees
// @Param id path int true "Employee ID"
// @Success 200 {string} string "OK"
// @Router /employees/{id} [delete]
func (i *EmployeeHandler) Delete(ctx echo.Context) error {
	traceID := trace_id.GetID(ctx)
	usecaseContext := trace_id.SetIDx(ctx.Request().Context(), traceID)

	ID, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
	if err != nil {
		return ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	i.Logger.InfoT(traceID, "delete employee by id", mlog.Any("id", ID))

	if err := i.EmployeeService.Delete(usecaseContext, ID); err != nil {
		return ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	} else {
		return SuccessResponse(ctx, http.StatusOK, nil)
	}
}
