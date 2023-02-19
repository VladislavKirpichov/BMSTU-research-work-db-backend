package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/usecase"
	"github.com/v.kirpichov/admin/pkg/errorHandler"
)

type EmployersHandler struct {
	usecase usecase.EmployerU
}

func NewEmployersHandler(usecase usecase.EmployerU) *EmployersHandler {
	return &EmployersHandler{
		usecase: usecase,
	}
}

type GetEmployerResponse struct {
	Employer *models.Employer `json:"employer"`
}

func (e *EmployersHandler) GetEmployer(c echo.Context) error {
	ctx := c.Request().Context()

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errorHandler.ErrClient
	}

	employer, err := e.usecase.GetEmployer(ctx, int64(id))
	if err != nil {
		return errorHandler.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, &GetEmployerResponse{
		Employer: employer,
	})

	return nil
}

type GetEmployersResponse struct {
	Employers []*models.Employer `json:"employers"`
}

func (e *EmployersHandler) GetEmployers(c echo.Context) error {
	ctx := c.Request().Context()

	employers, err := e.usecase.GetEmployers(ctx)
	if err != nil {
		return errorHandler.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, &GetEmployersResponse{
		Employers: employers,
	})

	return nil
}

type CreateEmployerRequest struct {
	Name string `json:"name"`
}

type CreateEmployerResponse struct {
	Id int64 `json:"id"`
}

func (e *EmployersHandler) CreateEmployer(c echo.Context) error {
	req := &CreateServiceRequest{}
	ctx := c.Request().Context()

	err := json.NewDecoder(c.Request().Body).Decode(req)
	if err != nil {
		return errorHandler.ErrClient
	}

	id, err := e.usecase.CreateEmployer(ctx, &models.Employer{
		Id:   0,
		Name: req.Name,
	})
	if err != nil {
		return errorHandler.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, &CreateEmployerResponse{
		Id: id,
	})

	return nil
}

func (e *EmployersHandler) UpdateEmployer(c echo.Context) error {
	req := &models.Employer{}
	ctx := c.Request().Context()

	err := json.NewDecoder(c.Request().Body).Decode(req)
	if err != nil {
		return errorHandler.ErrClient
	}

	err = e.usecase.UpdateEmployer(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployersHandler) DeleteEmployer(c echo.Context) error {
	ctx := c.Request().Context()

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errorHandler.ErrClient
	}

	err = e.usecase.DeleteEmployer(ctx, int64(id))
	if err != nil {
		return err
	}

	return nil
}
