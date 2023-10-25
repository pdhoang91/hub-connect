package mocks

import (
	"hub-connect/internal/entities"

	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (ur *UserRepository) Create(user *entities.User) error {
	ret := ur.Called(user)
	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.User) error); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}
	return r0
}

func (ur *UserRepository) Update(user *entities.User) error {
	ret := ur.Called(user)
	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.User) error); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}
	return r0
}

func (ur *UserRepository) FindByID(id int) (*entities.User, error) {
	ret := ur.Called(id)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(int) *entities.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
