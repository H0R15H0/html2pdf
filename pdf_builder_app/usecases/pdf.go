package usecases

import (
	"context"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
)

type PdfUsecaseConvertHtml2PdfCommand struct {
	UserID string
}

type PdfUsecase interface {
	ConvertHtml2Pdf(context.Context, PdfUsecaseConvertHtml2PdfCommand) (*entities.Pdf, error)
}

type pdfUsecase struct {
	pdfRepo  repositories.PdfRepo
	userRepo repositories.UserRepo
}

func NewPdfUsecase(up repositories.PdfRepo, u repositories.UserRepo) PdfUsecase {
	return &pdfUsecase{pdfRepo: up, userRepo: u}
}

func (u *pdfUsecase) ConvertHtml2Pdf(ctx context.Context, cmd PdfUsecaseConvertHtml2PdfCommand) (*entities.Pdf, error) {
	userID, err := values.UserIDString(cmd.UserID)
	if err != nil {
		return nil, err
	}

	pdf, err := u.pdfRepo.CreateWithSource(ctx, userID)
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
