package repositories

import (
	"context"
	"database/sql"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/infra/postgresql/models"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type partialPdfRepo struct {
	db *sql.DB
}

func NewPartialPdfRepo(db *sql.DB) repositories.PartialPdfRepo {
	return &partialPdfRepo{db: db}
}

func (r *partialPdfRepo) Create(ctx context.Context, unifiedPdfID values.PdfID, sourceHTMLUrl values.PartialPdfSourceHTMLUrl, number values.PartialPdfNumber, s3Url values.PartialPdfS3URL) (*entities.PartialPdf, error) {
	p := models.PartialPDF{ID: uuid.NewString(), UnifiedPDFID: unifiedPdfID.String(), SourceHTMLURL: string(sourceHTMLUrl), Number: int(number), S3URL: string(s3Url)}
	err := p.Insert(ctx, r.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	partialPdfEntity := entities.PartialPdf{
		ID:            values.MustPartialPdfIDString(p.ID),
		UnifiedPDFID:  values.MustPdfIDString(p.UnifiedPDFID),
		SourceHTMLURL: values.PartialPdfSourceHTMLUrl(p.SourceHTMLURL),
		Number:        values.PartialPdfNumber(p.Number),
		S3URL:         values.PartialPdfS3URL(p.S3URL),
	}
	return &partialPdfEntity, nil
}

func (r *partialPdfRepo) Update(ctx context.Context, pEntity *entities.PartialPdf) (*entities.PartialPdf, error) {
	p := models.PartialPDF{ID: pEntity.ID.String(), UnifiedPDFID: pEntity.UnifiedPDFID.String(), SourceHTMLURL: string(pEntity.SourceHTMLURL), Number: int(pEntity.Number), S3URL: string(pEntity.S3URL), PDFCreatedAt: null.Time(pEntity.PDFCreatedAt)}
	_, err := p.Update(ctx, r.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	partialPdfEntity := entities.PartialPdf{
		ID:            values.MustPartialPdfIDString(p.ID),
		UnifiedPDFID:  values.MustPdfIDString(p.UnifiedPDFID),
		SourceHTMLURL: values.PartialPdfSourceHTMLUrl(p.SourceHTMLURL),
		Number:        values.PartialPdfNumber(p.Number),
		S3URL:         values.PartialPdfS3URL(p.S3URL),
	}
	return &partialPdfEntity, nil
}
