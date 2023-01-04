package repositories

import (
	"context"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
)

type PdfRepo interface {
	CreateWithSource(context.Context, values.UserID) (*entities.Pdf, error)
	FindWithRelation(context.Context, values.UserID, values.PdfID) (*entities.Pdf, error)
}
