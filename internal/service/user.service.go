package service

import (
	"authentication/internal/apperrors"
	"authentication/internal/dto"
	"authentication/internal/models"
	"authentication/internal/repo"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindOne(dto dto.FindOne) (*models.User, error)
	CreateUser(dto dto.CreateUser) (*models.User, error)
	SignIn(user dto.CreateUser) (*models.User, error)
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
	user, err := u.dao.NewUserQuery().FindOneByEmail(*dto.Email)
	if user != nil {
		err = apperrors.AlreadyExistsErr
		return nil, err
	}

	return user, nil
}

func (u *userService) CreateUser(dto dto.CreateUser) (*models.User, error) {
	user, _ := u.dao.NewUserQuery().FindOneByEmail(*dto.Email)
	if user != nil {
		return nil, apperrors.AlreadyExistsErr
	}

	// hash user password
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), 12)
	if err != nil {
		return nil, err
	}

	if err := u.dao.NewUserQuery().InsertNewUser(*dto.Email, string(hashBytes)); err != nil {
		return nil, err
	}

	user, err = u.dao.NewUserQuery().FindOneByEmail(*dto.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) SignIn(dto dto.CreateUser) (*models.User, error) {

	user, err := u.dao.NewUserQuery().FindOneByEmail(*dto.Email)
	if err != nil {
		return nil, err
	}

	// compare stored hash with the given user password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*dto.Password))
	if err != nil {
		fmt.Println(err)
		return nil, apperrors.WrongPasswordErr
	}

	return user, err

}
