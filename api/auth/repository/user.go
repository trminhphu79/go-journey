package repository

import (
	"app/api/auth/model"
	"app/arch/postgres"
)

type IUserRepo interface {
	FindOne(id string) (*model.User, error)
}

type userRepo struct {
	db postgres.Database
}

func CreateUserRepository(db postgres.Database) IUserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) FindOne(id string) (*model.User, error) {
	return nil, nil
}
