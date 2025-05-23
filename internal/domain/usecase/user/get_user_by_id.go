package user

import (
	"github.com/0resama/go-clean-ddd/internal/domain/logger"
	"github.com/0resama/go-clean-ddd/internal/domain/model"
	"github.com/0resama/go-clean-ddd/internal/domain/repository"
)

type GetUserByIDUseCase struct {
	repo   repository.UserRepository
	logger logger.Logger
}

func NewGetUserByIDUseCase(r repository.UserRepository, l logger.Logger) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{repo: r, logger: l}
}

func (uc *GetUserByIDUseCase) Execute(id uint) (*model.User, error) {
	return uc.repo.GetByID(id)
}
