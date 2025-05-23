package user

import (
	"github.com/0resama/go-clean-ddd/internal/domain/logger"
	"github.com/0resama/go-clean-ddd/internal/domain/model"
	"github.com/0resama/go-clean-ddd/internal/domain/repository"
)

type CreateUser struct {
	repo   repository.UserRepository
	logger logger.Logger
}

func NewCreateUser(r repository.UserRepository, l logger.Logger) *CreateUser {
	return &CreateUser{repo: r, logger: l}
}

func (uc *CreateUser) Execute(user *model.User) error {
	uc.logger.Info("Creating user", "email", user.Email)
	return uc.repo.Create(user)
}
