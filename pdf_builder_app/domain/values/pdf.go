package values

import "github.com/google/uuid"

type PdfID uuid.UUID

var ZeroPdfID PdfID

func (u PdfID) String() string {
	return uuid.UUID(u).String()
}

func PdfIDString(s string) (PdfID, error) {
	pdfId, err := uuid.Parse(s)
	if err != nil {
		return ZeroPdfID, err
	}
	return PdfID(pdfId), nil
}

func MustPdfIDString(s string) PdfID {
	pdfId, err := PdfIDString(s)
	if err != nil {
		panic("PdfID must be valid")
	}
	return pdfId
}

type PdfS3URL string
