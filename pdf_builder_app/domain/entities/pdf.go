package entities

import (
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
	"github.com/volatiletech/null/v8"
)

type Pdf struct {
	ID         values.PdfID    `json:"id"`
	UserID     values.UserID   `json:"userId"`
	S3URL      values.PdfS3URL `json:"s3Url"`
	UnifiedAt  null.Time       `json:"unifiedAt"`
	PartialPdf []*PartialPdf   `json:"partialPdf"`
}
