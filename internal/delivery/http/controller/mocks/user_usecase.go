package mocks

import (
	"hub-connect/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

type UserUseCase struct {
	mock.Mock
}

// CreateUser handles the creation of a new user.
func (_m *UserUseCase) CreateUser(name string, email string) (*entities.User, error) {

	ret := _m.Called(name, email)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(string, string) *entities.User); ok {
		r0 = rf(name, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser updates a user's associated team.
func (_m *UserUseCase) UpdateUser(id, teamID int) (*entities.User, error) {
	return nil, nil
}

// GetUserByID retrieves user information by its ID.
func (_m *UserUseCase) GetUserByID(id int) (*entities.User, error) {
	return nil, nil
}

// UserJoinTeam associates a user with a team.
func (_m *UserUseCase) UserJoinTeam(userID, teamID int) error {
	return nil
}
