package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/usecase"
	"github.com/v.kirpichov/admin/pkg/errorHandler"
)

type ServicesHandler struct {
	usecase usecase.ServicesU
}

func NewServicesUsecase(usecase usecase.ServicesU) *ServicesHandler {
	return &ServicesHandler{
		usecase: usecase,
	}
}

type GetServiceResponse struct {
	Service *models.Service `json:"service"`
}

func (s *ServicesHandler) GetService(c echo.Context) error {
	ctx := c.Request().Context()

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errorHandler.ErrClient
	}

	service, err := s.usecase.GetService(ctx, int64(id))
	if err != nil {
		return errorHandler.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, &GetServiceResponse{
		Service: service,
	})

	return nil
}

type CreateServiceRequest struct {
	Name string `json:"name"`
}

type CreateServiceResponse struct {
	Id int64 `json:"id"`
}

func (s *ServicesHandler) CreateService(c echo.Context) error {
	req := &CreateServiceRequest{}
	ctx := c.Request().Context()

	fmt.Println("before json")

	err := json.NewDecoder(c.Request().Body).Decode(req)
	if err != nil {
		return errorHandler.ErrClient
	}

	id, err := s.usecase.CreateService(ctx, &models.Service{
		Id:   0,
		Name: req.Name,
	})
	if err != nil {
		return errorHandler.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, &CreateServiceResponse{
		Id: id,
	})

	return nil
}

func (s *ServicesHandler) UpdateService(c echo.Context) error {
	req := &models.Service{}
	ctx := c.Request().Context()

	err := json.NewDecoder(c.Request().Body).Decode(req)
	if err != nil {
		return errorHandler.ErrClient
	}

	err = s.usecase.UpdateService(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServicesHandler) DeleteService(c echo.Context) error {
	ctx := c.Request().Context()

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errorHandler.ErrClient
	}

	err = s.usecase.DeleteService(ctx, int64(id))
	if err != nil {
		return err
	}

	return nil
}
