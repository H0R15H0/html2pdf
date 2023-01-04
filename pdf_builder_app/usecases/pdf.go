package usecases

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
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
	Unify(context.Context, PdfUsecaseUnifyCommand, http.ResponseWriter) (*string, error)
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

func (u *pdfUsecase) Unify(ctx context.Context, cmd PdfUsecaseUnifyCommand, w http.ResponseWriter) (*string, error) {
	userID, err := values.UserIDString(cmd.UserID)
	if err != nil {
		return nil, err
	}
	pdfID, err := values.PdfIDString(cmd.PdfID)
	if err != nil {
		return nil, err
	}

	pdf, err := u.pdfRepo.FindWithRelation(ctx, userID, pdfID)
	if err != nil {
		return nil, err
	}

	var partialPdfs []io.ReadSeeker
	for _, pPdf := range pdf.PartialPdfs {
		body, err := u.filePartialRepo.GetObject(ctx, values.FilePdfKey(pPdf.ID.String()))
		if err != nil {
			return nil, err
		}
		n, err := io.ReadAll(body)
		if err != nil {
			return nil, err
		}
		partialPdfs = append(partialPdfs, bytes.NewReader(n))
	}

	config := pdfcpu.NewDefaultConfiguration()
	err = api.Merge(partialPdfs, w, config)

	pdfName := fmt.Sprintf("%s.pdf", pdfID.String())

	return &pdfName, nil
}
