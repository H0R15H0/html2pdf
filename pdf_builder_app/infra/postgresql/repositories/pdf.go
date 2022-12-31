package repositories

import (
	"context"
	"database/sql"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/infra/postgresql/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type pdfRepo struct {
	db *sql.DB
}

func NewPdfRepo(db *sql.DB) repositories.PdfRepo {
	return &pdfRepo{db: db}
}

func (r *pdfRepo) CreateWithSource(ctx context.Context, userID values.UserID) (*entities.Pdf, error) {
	// TODO: generate s3_url
	pdf := models.UsersPDF{ID: uuid.NewString(), UserID: userID.String(), S3URL: "s3_url"}
	err := pdf.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	pdfEntity := entities.Pdf{
		ID:     values.MustPdfIDString(pdf.ID),
		UserID: userID,
		S3URL:  values.PdfS3URL(pdf.S3URL),
	}
	return &pdfEntity, nil
}
