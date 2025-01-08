package storage

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserStorage struct {
	mock.Mock
}

func (m *MockUserStorage) CreateUser(user User) (User, error) {
	args := m.Called(user)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockUserStorage) GetUserByID(id uuid.UUID) (User, error) {
	args := m.Called(id)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockUserStorage) UpdateUser(id uuid.UUID, user User) (User, error) {
	args := m.Called(id, user)
	return args.Get(0).(User), args.Error(1)
}
