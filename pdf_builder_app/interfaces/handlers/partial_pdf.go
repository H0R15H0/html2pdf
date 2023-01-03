package handlers

import (
	"fmt"
	"net/http"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/usecases"
	"github.com/labstack/echo/v4"
)

type PartialPdfHandler interface {
	Create(echo.Context) error
	Error(echo.Context) error
}

type partialPdfHandler struct {
	partialPdfUsecase usecases.PartialPdfUsecase
}

func NewPartialPdfHandler(partialPdfUsecase usecases.PartialPdfUsecase) PartialPdfHandler {
	return &partialPdfHandler{
		partialPdfUsecase: partialPdfUsecase,
	}
}

func (u *partialPdfHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var cmd usecases.PartialPdfUsecaseConvertHtml2PdfCommand
	err := c.Bind(&cmd)
	if err != nil {
		return JsonError(c, err, &APIError{
			Message: "リクエストが正しくありません。",
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	pdf, err := u.partialPdfUsecase.ConvertHtml2Pdf(ctx, cmd)
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

func (u *partialPdfHandler) Error(c echo.Context) error {
	fmt.Println("html2pdf services convert failed.")
	return c.JSON(http.StatusOK, "")
}
