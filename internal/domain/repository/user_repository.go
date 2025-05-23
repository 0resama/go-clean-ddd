package repository

import "github.com/0resama/go-clean-ddd/internal/domain/model"

type UserRepository interface {
    Create(user *model.User) error
    GetByID(id uint) (*model.User, error)
}
