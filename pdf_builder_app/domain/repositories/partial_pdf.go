package repositories

import (
	"context"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
)

type PartialPdfRepo interface {
	Create(context.Context, values.PdfID, values.PartialPdfSourceHTMLUrl, values.PartialPdfNumber, values.PartialPdfS3URL) (*entities.PartialPdf, error)
	Update(context.Context, *entities.PartialPdf) (*entities.PartialPdf, error)
}
