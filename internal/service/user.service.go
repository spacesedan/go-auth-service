package service

import (
	"authentication/internal/apperrors"
	"authentication/internal/dto"
	"authentication/internal/models"
	"authentication/internal/repo"
	"fmt"
)

type UserService interface {
	FindOne(dto dto.FindOne) (*models.User, error)
	CreateUser(dto dto.CreateUser) (*models.User, error)
}

type userService struct {
	dao repo.DAO
}

func NewUserService(dao repo.DAO) UserService {
	return &userService{
		dao: dao,
	}
}

func (u *userService) FindOne(dto dto.FindOne) (*models.User, error) {
	email := *dto.Email

	user, err := u.dao.NewUserQuery().FindOneByEmail(email)
	if user != nil {
		err = apperrors.AlreadyExists()
		return nil, err
	}

	return user, nil
}

func (u *userService) CreateUser(dto dto.CreateUser) (*models.User, error) {
	user, _ := u.dao.NewUserQuery().FindOneByEmail(*dto.Email)
	if user != nil {
		return nil, apperrors.AlreadyExists()
	}

	fmt.Println(*dto.Email, *dto.Password)
	user, err := u.dao.NewUserQuery().InsertNewUser(*dto.Email, *dto.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
