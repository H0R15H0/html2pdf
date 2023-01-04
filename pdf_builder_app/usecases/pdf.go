package usecases

import (
	"context"
	"io"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
)

type PdfUsecaseCreateCommand struct {
	UserID string
}

type PdfUsecaseUnifyCommand struct {
	UserID string
	PdfID  string
}

type PdfUsecase interface {
	Create(context.Context, PdfUsecaseCreateCommand) (*entities.Pdf, error)
	Unify(context.Context, PdfUsecaseUnifyCommand) (*io.Reader, error)
}

type pdfUsecase struct {
	pdfRepo         repositories.PdfRepo
	userRepo        repositories.UserRepo
	filePartialRepo repositories.FilePartialPdfRepo
}

func NewPdfUsecase(up repositories.PdfRepo, u repositories.UserRepo, f repositories.FilePartialPdfRepo) PdfUsecase {
	return &pdfUsecase{pdfRepo: up, userRepo: u, filePartialRepo: f}
}

func (u *pdfUsecase) Create(ctx context.Context, cmd PdfUsecaseCreateCommand) (*entities.Pdf, error) {
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

func (u *pdfUsecase) Unify(ctx context.Context, cmd PdfUsecaseUnifyCommand) (*io.Reader, error) {
	userID, err := values.UserIDString(cmd.UserID)
	if err != nil {
		return nil, err
	}
	pdfID, err := values.PdfIDString(cmd.PdfID)
	if err != nil {
		return nil, err
	}

	_, err = u.pdfRepo.FindWithRelation(ctx, userID, pdfID)
	if err != nil {
		return nil, err
	}

	return nil, nil

	// var r io.Reader
	// for _, pPdf := range pdf.PartialPdf {
	// 	body, err := u.filePartialRepo.GetObject(ctx, values.FilePdfKey(pPdf.S3URL))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	r = body
	// }

	// return &r, nil
}
