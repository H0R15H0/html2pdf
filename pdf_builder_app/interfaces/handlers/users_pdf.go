package handlers

import (
	"net/http"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/usecases"
	"github.com/labstack/echo/v4"
)

type UsersPdfHandler interface {
	InitializeUsersPdf(echo.Context) error
	Unify(echo.Context) error
}

type usersPdfHandler struct {
	pdfUsecase usecases.PdfUsecase
}

func NewUsersPdfHandler(pdfUsecase usecases.PdfUsecase) UsersPdfHandler {
	return &usersPdfHandler{
		pdfUsecase: pdfUsecase,
	}
}

func (u *usersPdfHandler) InitializeUsersPdf(c echo.Context) error {
	ctx := c.Request().Context()

	cmd := usecases.PdfUsecaseCreateCommand{UserID: c.Param("id")}

	pdf, err := u.pdfUsecase.Create(ctx, cmd)
	if err != nil {
		return JsonError(c, err, &APIError{
			Message: "ユーザーが見つかりませんでした",
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	data := map[string]*entities.Pdf{"pdf": pdf}
	return c.JSON(http.StatusOK, data)
}

func (u *usersPdfHandler) Unify(c echo.Context) error {
	ctx := c.Request().Context()

	cmd := usecases.PdfUsecaseUnifyCommand{UserID: c.Param("user_id"), PdfID: c.Param("pdf_id")}

	resp := c.Response()
	pdfName, err := u.pdfUsecase.Unify(ctx, cmd, resp.Writer)
	if err != nil {
		return JsonError(c, err, &APIError{
			Message: "PDFが見つかりませんでした",
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	resp.Header().Set("Cache-Control", "no-store")
	resp.Header().Set(echo.HeaderContentType, "application/pdf")
	resp.Header().Set(echo.HeaderAccessControlExposeHeaders, "Content-Disposition")
	resp.Header().Set(echo.HeaderContentDisposition, "attachment; filename="+*pdfName)

	return c.NoContent(http.StatusOK)
}
