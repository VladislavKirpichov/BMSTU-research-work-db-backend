package errorHandler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO: переделать ошибки (Handler по сути конвертит обычные ошибки в echo.HTTPError)
var (
	ErrInternal      = NewInternalServerError("undefined error")
	ErrSecret        = NewInternalSecretError("secret")
	ErrNoUser        = echo.NewHTTPError(http.StatusNotFound, "No such user")
	ErrWrongPassword = NewClientError("Wrong password")
	ErrClient        = echo.NewHTTPError(http.StatusNotFound, "Client error")
	ErrUnauthorized  = NewClientError("Unauthorized")

	ErrInvalidSession = echo.NewHTTPError(http.StatusUnauthorized, "provide valid credentials")
)

type ErrorHandler struct {
	statusCodes map[error]int
}

func New() *ErrorHandler {
	return &ErrorHandler{
		statusCodes: map[error]int{
			ErrInternal:       http.StatusInternalServerError,
			ErrSecret:         http.StatusInternalServerError,
			ErrNoUser:         http.StatusNotFound,
			ErrWrongPassword:  http.StatusBadRequest,
			ErrClient:         http.StatusBadRequest,
			ErrUnauthorized:   http.StatusUnauthorized,
			ErrInvalidSession: http.StatusUnauthorized,
		},
	}
}

func (e *ErrorHandler) getStatusCode(err error) int {
	code, exists := e.statusCodes[err]
	if !exists {
		return http.StatusInternalServerError
	}

	return code
}

func (e *ErrorHandler) Handler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = &echo.HTTPError{
			Code:    e.getStatusCode(he),
			Message: err.Error(),
		}
	}

	fmt.Println(err)

	if !c.Response().Committed {
		message := he.Message
		if _, ok := he.Message.(string); ok {
			message = map[string]interface{}{
				"message": err.Error(),
			}
		}

		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(he.Code, message)
		}

		if err != nil {
			c.Logger().Error(err)
		}
	}
}
