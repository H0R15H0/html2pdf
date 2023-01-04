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
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (r *pdfRepo) FindWithRelation(ctx context.Context, userID values.UserID, pdfID values.PdfID) (*entities.Pdf, error) {
	pdf, err := models.UsersPDFS(
		models.UsersPDFWhere.UserID.EQ(userID.String()),
		models.UsersPDFWhere.ID.EQ(pdfID.String()),
		qm.Load(qm.Rels(models.UsersPDFRels.UnifiedPDFPartialPDFS)),
	).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	var partialPdfs []*entities.PartialPdf

	for _, pp := range pdf.R.UnifiedPDFPartialPDFS {
		partialPdfs = append(partialPdfs, &entities.PartialPdf{
			ID:            values.MustPartialPdfIDString(pp.ID),
			UnifiedPDFID:  values.MustPdfIDString(pp.UnifiedPDFID),
			SourceHTMLURL: values.PartialPdfSourceHTMLUrl(pp.SourceHTMLURL),
			Number:        values.PartialPdfNumber(pp.Number),
			S3URL:         values.PartialPdfS3URL(pp.S3URL),
			PDFCreatedAt:  values.PartialPdfCreatedAt(pp.PDFCreatedAt),
		})
	}

	pdfEntity := entities.Pdf{
		ID:          values.MustPdfIDString(pdf.ID),
		UserID:      userID,
		S3URL:       values.PdfS3URL(pdf.S3URL),
		PartialPdfs: partialPdfs,
	}

	return &pdfEntity, nil
}
