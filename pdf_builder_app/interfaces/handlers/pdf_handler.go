package handlers

import "github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"

type PdfHandler interface{}

type pdfHandler struct {
	pdfRepo repositories.PdfRepo
}

func NewPdfHandler(pdfRepo repositories.PdfRepo) PdfHandler {
	return &pdfHandler{
		pdfRepo: pdfRepo,
	}
}
