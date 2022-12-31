package repositories

import (
	"context"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
)

type UserRepo interface {
	FindUser(context.Context, string) (*entities.User, error)
}
