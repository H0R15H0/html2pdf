package usecases

import (
	"context"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/entities"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
	"github.com/google/uuid"
)

type UserUsecaseGetUserCommand struct {
	ID string
}

type UserUsecaseCreateUserCommand struct{}

type UserUsecase interface {
	GetUser(context.Context, UserUsecaseGetUserCommand) (*entities.User, error)
	CreateUser(context.Context, UserUsecaseCreateUserCommand) (*entities.User, error)
}

type userUsecase struct {
	userRepo repositories.UserRepo
}

func NewUserUsecase(u repositories.UserRepo) UserUsecase {
	return &userUsecase{userRepo: u}
}

func (u *userUsecase) GetUser(ctx context.Context, cmd UserUsecaseGetUserCommand) (*entities.User, error) {
	id := cmd.ID
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	user, err := u.userRepo.FindUser(ctx, values.UserID(uuid.String()))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, cmd UserUsecaseCreateUserCommand) (*entities.User, error) {
	user, err := u.userRepo.CreateUser(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
