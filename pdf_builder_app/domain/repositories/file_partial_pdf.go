package repositories

import (
	"context"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
)

type FilePartialPdfRepo interface {
	CreatePreSignedUrl(context.Context, values.FilePdfKey) (values.FilePreSignedUrl, error)
}
