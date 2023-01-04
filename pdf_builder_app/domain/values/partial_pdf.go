package values

import (
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
)

type PartialPdfID string

var ZeroPartialPdfID PartialPdfID

func (u PartialPdfID) String() string {
	return string(u)
}

func PartialPdfIDString(s string) (PartialPdfID, error) {
	pdfId, err := uuid.Parse(s)
	if err != nil {
		return ZeroPartialPdfID, err
	}
	return PartialPdfID(pdfId.String()), nil
}

func MustPartialPdfIDString(s string) PartialPdfID {
	pdfId, err := PartialPdfIDString(s)
	if err != nil {
		panic("PartialPdfID must be valid")
	}
	return pdfId
}

type PartialPdfSourceHTMLUrl string
type PartialPdfS3URL string
type PartialPdfNumber int
type PartialPdfCreatedAt null.Time
