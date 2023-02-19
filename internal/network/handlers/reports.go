package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/usecase"
	"github.com/v.kirpichov/admin/pkg/errorHandler"
)

type ReportsHandler struct {
	usecase usecase.ReportU
}

func NewReportsHandler(usecase usecase.ReportU) *ReportsHandler {
	return &ReportsHandler{
		usecase: usecase,
	}
}

type GetReportResponse struct {
	Report *models.Report `json:"report"`
}

func (h *ReportsHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errorHandler.ErrClient
	}

	report, err := h.usecase.Get(ctx, int64(id))
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, &GetReportResponse{
		Report: report,
	})
	return nil
}

type GetReportsReponse struct {
	Reports []*models.Report `json:"reports"`
}

func (h *ReportsHandler) GetReports(c echo.Context) error {
	ctx := c.Request().Context()

	reports, err := h.usecase.GetReports(ctx)
	if err != nil {
		return errorHandler.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, &GetReportsReponse{
		Reports: reports,
	})

	return nil
}

type CreateReportResponse struct {
	Id int64 `json:"id"`
}

func (h *ReportsHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := h.usecase.Create(ctx)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, &CreateReportResponse{
		Id: id,
	})

	return nil
}
