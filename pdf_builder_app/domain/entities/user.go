package entities

import (
	"time"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
)

type User struct {
	ID        values.UserID `json:"id"`
	CreatedAt time.Time     `json:"createdAt"`
}
