package usecases

import (
	"context"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
)

type UsersPdfUsecaseInitializeUsersPdfCommand struct {
	UserID string
}

type UsersPdfUsecase interface {
	InitializeUsersPdf(context.Context, UsersPdfUsecaseInitializeUsersPdfCommand) (*entities.Pdf, error)
}

type usersPdfUsecase struct {
	pdfRepo  repositories.PdfRepo
	userRepo repositories.UserRepo
}

func NewUsersPdfUsecase(up repositories.PdfRepo, u repositories.UserRepo) UsersPdfUsecase {
	return &usersPdfUsecase{pdfRepo: up, userRepo: u}
}

func (u *usersPdfUsecase) InitializeUsersPdf(ctx context.Context, cmd UsersPdfUsecaseInitializeUsersPdfCommand) (*entities.Pdf, error) {
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
