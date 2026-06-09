package handler

import (
	"echo-practice/internal/model"
	"echo-practice/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		svc: s,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.svc.CreateUser(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.svc.GetUserProfile(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if err := h.svc.DeleteUser(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}
