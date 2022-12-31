package handlers

import (
	"net/http"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/usecases"
	"github.com/labstack/echo/v4"
)

type UsersPdfHandler interface {
	InitializeUsersPdf(echo.Context) error
}

type usersPdfHandler struct {
	usersPdfUsecase usecases.UsersPdfUsecase
}

func NewUsersPdfHandler(usersPdfUsecase usecases.UsersPdfUsecase) UsersPdfHandler {
	return &usersPdfHandler{
		usersPdfUsecase: usersPdfUsecase,
	}
}

func (u *usersPdfHandler) InitializeUsersPdf(c echo.Context) error {
	ctx := c.Request().Context()

	cmd := usecases.UsersPdfUsecaseInitializeUsersPdfCommand{UserID: c.Param("id")}

	pdf, err := u.usersPdfUsecase.InitializeUsersPdf(ctx, cmd)
	if err != nil {
		return JsonError(c, err, &APIError{
			Message: "ユーザーが見つかりませんでした",
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	data := map[string]interface{}{"pdf": pdf}
	return c.JSON(http.StatusOK, data)
}
