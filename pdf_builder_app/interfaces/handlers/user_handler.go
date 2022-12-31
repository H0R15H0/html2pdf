package handlers

import (
	"net/http"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUser(echo.Context) error
}

type userHandler struct {
	userRepo repositories.UserRepo
}

func NewUserHandler(userRepo repositories.UserRepo) UserHandler {
	return &userHandler{
		userRepo: userRepo,
	}
}

func (u *userHandler) GetUser(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()

	user, err := u.userRepo.FindUser(ctx, id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, user.ID)
}
