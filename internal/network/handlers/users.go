package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/siruspen/logrus"
	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/usecase"
)

type UserHandler struct {
	cfg     *configs.Config
	usecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase, cfg *configs.Config) *UserHandler {
	return &UserHandler{
		cfg:     cfg,
		usecase: usecase,
	}
}

func (u *UserHandler) SignIn(c echo.Context) error {
	return nil
}

func (u *UserHandler) SignUp(c echo.Context) error {
	var input = models.InputUser{}
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		c.Error(err)
		return err
	}

	userId, err := u.usecase.SignUp(c.Request().Context(), &input)
	if err != nil {
		c.Error(err)
		return err
	}

	c.String(http.StatusOK, strconv.Itoa(int(userId)))
	return nil
}

type UsersResponse struct {
	Users []*models.User
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := u.usecase.GetAllUsers(c.Request().Context())
	if err != nil {
		logrus.Error(err)
		c.Error(err)
		return err
	}

	c.JSON(http.StatusOK, UsersResponse{
		Users: users,
	})

	return nil
}
