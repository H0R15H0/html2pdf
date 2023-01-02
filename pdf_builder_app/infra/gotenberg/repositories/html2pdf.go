package repositories

import (
	"context"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/infra/gotenberg"
)

type html2PdfRepo struct {
	client *gotenberg.GotenbergClient
}

func NewHtml2PdfRepo(c *gotenberg.GotenbergClient) repositories.Html2PdfRepo {
	return &html2PdfRepo{client: c}
}

func (r *html2PdfRepo) Send(ctx context.Context, htmlUrl values.Html2PdfHtmlUrl, s3Url values.FilePreSignedUrl) error {
	err := r.client.ConvertURLWithWebhook(ctx, string(htmlUrl), string(s3Url))
	return err
}
