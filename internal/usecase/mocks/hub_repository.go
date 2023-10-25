package mocks

import (
	"hub-connect/internal/entities"

	"github.com/stretchr/testify/mock"
)

type HubRepository struct {
	mock.Mock
}

func (hr *HubRepository) Create(hub *entities.Hub) error {
	ret := hr.Called(hub)
	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Hub) error); ok {
		r0 = rf(hub)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}
	return r0
}

func (hr *HubRepository) FindByID(id int) (*entities.Hub, error) {
	ret := hr.Called(id)

	var r0 *entities.Hub
	if rf, ok := ret.Get(0).(func(int) *entities.Hub); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Hub)
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

func (hr *HubRepository) SearchHubs(keyword string) ([]*entities.Hub, error) {
	ret := hr.Called(keyword)

	var r0 []*entities.Hub
	if rf, ok := ret.Get(0).(func(string) []*entities.Hub); ok {
		r0 = rf(keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Hub)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(keyword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
