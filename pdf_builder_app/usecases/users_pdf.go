package usecases

import (
	"context"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
)

type UsersPdfUsecaseConvertHtml2PdfCommand struct {
	UserID string
}

type UsersPdfUsecase interface {
	ConvertHtml2Pdf(context.Context, UsersPdfUsecaseConvertHtml2PdfCommand) (*entities.Pdf, error)
}

type usersPdfUsecase struct {
	pdfRepo  repositories.PdfRepo
	userRepo repositories.UserRepo
}

func NewUsersPdfUsecase(up repositories.PdfRepo, u repositories.UserRepo) UsersPdfUsecase {
	return &usersPdfUsecase{pdfRepo: up, userRepo: u}
}

func (u *usersPdfUsecase) ConvertHtml2Pdf(ctx context.Context, cmd UsersPdfUsecaseConvertHtml2PdfCommand) (*entities.Pdf, error) {
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
