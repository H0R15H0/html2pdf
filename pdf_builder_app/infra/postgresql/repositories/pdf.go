package repositories

import (
	"database/sql"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
)

type pdfRepo struct {
	db *sql.DB
}

func NewPdfRepo(db *sql.DB) repositories.PdfRepo {
	return &pdfRepo{db: db}
}
