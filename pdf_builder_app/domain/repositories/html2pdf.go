package repositories

import (
	"context"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
)

type Html2PdfRepo interface {
	Send(context.Context, values.Html2PdfHtmlUrl) error
}
