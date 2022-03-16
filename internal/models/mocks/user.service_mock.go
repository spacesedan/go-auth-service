package mocks

import (
	"authentication/internal/dto"
	"authentication/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockedUserService struct {
	mock.Mock
}

func (m *MockedUserService) FindOne(findOne dto.FindOne) (*models.User, error) {
	ret := m.Called(findOne)

	var r0 *models.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.User)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockedUserService) CreateUser(dto dto.CreateUser) (*models.User, error) {
	ret := m.Called(dto)

	var r0 *models.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.User)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
