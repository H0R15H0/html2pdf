package usecases

import (
	"context"
	"time"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
	"github.com/volatiletech/null/v8"
)

type PartialPdfUsecaseConvertHtml2PdfCommand struct {
	UnifiedPdfID  string `param:"pdf_id"`
	SourceHTMLUrl string `json:"sourceHtmlUrl"`
	Number        int    `json:"number"`
}

type PartialPdfUsecase interface {
	ConvertHtml2Pdf(context.Context, PartialPdfUsecaseConvertHtml2PdfCommand) (*entities.PartialPdf, error)
}

type partialPdfUsecase struct {
	partialPdfRepo     repositories.PartialPdfRepo
	html2PdfRepo       repositories.Html2PdfRepo
	filePartialPdfRepo repositories.FilePartialPdfRepo
}

func NewPartialPdfUsecase(pp repositories.PartialPdfRepo, h repositories.Html2PdfRepo, f repositories.FilePartialPdfRepo) PartialPdfUsecase {
	return &partialPdfUsecase{partialPdfRepo: pp, html2PdfRepo: h, filePartialPdfRepo: f}
}

func (u *partialPdfUsecase) ConvertHtml2Pdf(ctx context.Context, cmd PartialPdfUsecaseConvertHtml2PdfCommand) (*entities.PartialPdf, error) {
	uPdfID, err := values.PdfIDString(cmd.UnifiedPdfID)
	if err != nil {
		return nil, err
	}

	pPdf, err := u.partialPdfRepo.Create(ctx, uPdfID, values.PartialPdfSourceHTMLUrl(cmd.SourceHTMLUrl), values.PartialPdfNumber(cmd.Number), values.PartialPdfS3URL("s3_url"))
	if err != nil {
		return nil, err
	}

	preSignedUrl, err := u.filePartialPdfRepo.CreatePreSignedUrl(ctx, values.FilePdfKey(pPdf.ID.String()))
	if err != nil {
		return nil, err
	}

	err = u.html2PdfRepo.Send(ctx, values.Html2PdfHtmlUrl(pPdf.SourceHTMLURL), preSignedUrl)
	if err != nil {
		return nil, err
	}

	pPdf.PDFCreatedAt = values.PartialPdfCreatedAt(null.NewTime(time.Now(), true))

	pPdf, err = u.partialPdfRepo.Update(ctx, pPdf)
	if err != nil {
		return nil, err
	}

	return pPdf, nil
}
