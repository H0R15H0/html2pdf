package entities

import (
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
)

type PartialPdf struct {
	ID            values.PartialPdfID            `json:"id"`
	UnifiedPDFID  values.PdfID                   `json:"unifiedPdfId"`
	SourceHTMLURL values.PartialPdfSourceHTMLUrl `json:"sourceHtmlUrl"`
	Number        values.PartialPdfNumber        `json:"number"`
	S3URL         values.PartialPdfS3URL         `json:"s3Url"`
	PDFCreatedAt  values.PartialPdfCreatedAt     `json:"pdfCreatedAt"`
}
