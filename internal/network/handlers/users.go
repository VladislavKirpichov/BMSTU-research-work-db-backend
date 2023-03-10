package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/siruspen/logrus"
	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/usecase"
	"github.com/v.kirpichov/admin/pkg/errorHandler"
)

type UserHandler struct {
	cfg     *configs.Config
	usecase usecase.UserU
}

func NewUserHandler(usecase usecase.UserU, cfg *configs.Config) *UserHandler {
	return &UserHandler{
		cfg:     cfg,
		usecase: usecase,
	}
}

type SignInResponse struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (u *UserHandler) SignIn(c echo.Context) error {
	var input = models.SignInUser{}
	ctx := c.Request().Context()

	sessionToken, err := u.usecase.GetSessionToken(ctx, input.Email)
	if err != nil {
		c.Error(err)
		return err
	}

	err = json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		return errorHandler.ErrClient
	}

	user, err := u.usecase.SignIn(ctx, &input)
	if err != nil {
		return err
	}

	sessionCookie := &http.Cookie{
		Name:     "session",
		Value:    sessionToken,
		Expires:  time.Now().Add(time.Duration(u.cfg.SessionConfig.ExpiresAt * int(time.Hour.Nanoseconds()))),
		Secure:   true,
		HttpOnly: true,
	}

	c.SetCookie(sessionCookie)
	c.JSON(http.StatusOK, SignInResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})

	return nil
}

type SignUpResponse struct {
	Id int64 `json:"id"`
}

func (u *UserHandler) SignUp(c echo.Context) error {
	var input = models.InputUser{}

	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		c.Error(err)
		return err
	}

	sessinToken, err := u.usecase.GetSessionToken(c.Request().Context(), input.Email)
	if err != nil {
		c.Error(err)
		return err
	}

	userId, err := u.usecase.SignUp(c.Request().Context(), &input)
	if err != nil {
		c.Error(err)
		return err
	}

	sessionCookie := &http.Cookie{
		Name:     "session",
		Value:    sessinToken,
		Expires:  time.Now().Add(time.Duration(u.cfg.SessionConfig.ExpiresAt * int(time.Hour.Nanoseconds()))),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}

	c.SetCookie(sessionCookie)
	c.JSON(http.StatusOK, &SignUpResponse{
		Id: userId,
	})
	return nil
}

type AuthResponse struct {
	Profile *models.User `json:"profile"`
}

func (u *UserHandler) Auth(c echo.Context) error {
	cookies := c.Cookies()
	ctx := c.Request().Context()

	for _, cookie := range cookies {
		if cookie.Name == "session" {
			user, err := u.usecase.Auth(ctx, cookie.Value)
			if err != nil {
				return errorHandler.ErrInvalidSession
			}

			c.JSON(http.StatusOK, &AuthResponse{
				Profile: user,
			})

			return nil
		}
	}

	return errorHandler.ErrInvalidSession
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

// TODO: logout
