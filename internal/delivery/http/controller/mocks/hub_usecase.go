package mocks

import (
	"hub-connect/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

type HubUseCase struct {
	mock.Mock
}

// CreateHub handles the creation of a new hub.
func (_m *HubUseCase) CreateHub(name, location string) (*entities.Hub, error) {

	ret := _m.Called(name, location)

	var r0 *entities.Hub
	if rf, ok := ret.Get(0).(func(string, string) *entities.Hub); ok {
		r0 = rf(name, location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Hub)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHubByID retrieves hub information by its ID.
func (hi *HubUseCase) GetHubByID(id int) (*entities.Hub, error) {
	return nil, nil
}

// SearchHubs searches for hubs based on a keyword.
func (hi *HubUseCase) SearchHubs(keyword string) ([]*entities.Hub, error) {
	return nil, nil
}
