package handlers

import (
	"net/http"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/usecases"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUser(echo.Context) error
	CreateUser(echo.Context) error
}

type userHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(userUsecase usecases.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (u *userHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	cmd := usecases.UserUsecaseGetUserCommand{ID: id}

	user, err := u.userUsecase.GetUser(ctx, cmd)
	if err != nil {
		return JsonError(c, err, &APIError{
			Message: "ユーザーが見つかりませんでした",
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *userHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := u.userUsecase.CreateUser(ctx, usecases.UserUsecaseCreateUserCommand{})
	if err != nil {
		return JsonError(c, err, &APIError{
			Message: "ユーザーが作成できませんでした",
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeServerHoge),
		})
	}

	return c.JSON(http.StatusOK, user)
}
