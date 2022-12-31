package repositories

import (
	"context"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
)

type UserRepo interface {
	FindUser(context.Context, values.UserID) (*entities.User, error)
	CreateUser(context.Context) (*entities.User, error)
}
