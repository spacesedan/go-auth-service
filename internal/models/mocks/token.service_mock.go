package mocks

import (
	"authentication/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockedTokenService struct {
	mock.Mock
}

func (m *MockedTokenService) GenerateJWT(user *models.User) (string, error) {
	ret := m.Called(user)

	var r0 string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockedTokenService) ParseJWT(userToken string) (*models.User, error) {
	ret := m.Called(userToken)

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
