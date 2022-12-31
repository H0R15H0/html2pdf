package repositories

import (
	"context"
	"database/sql"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/infra/postgresql/models"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repositories.UserRepo {
	return &userRepo{db: db}
}

func (u *userRepo) FindUser(ctx context.Context, id string) (*entities.User, error) {
	user, err := models.Users(models.UserWhere.ID.EQ(id)).One(ctx, u.db)
	if err != nil {
		return nil, err
	}
	user_entity := entities.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
	}
	return &user_entity, nil
}
