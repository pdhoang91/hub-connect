package mocks

import (
	"hub-connect/internal/entities"

	"github.com/stretchr/testify/mock"
)

type TeamRepository struct {
	mock.Mock
}

func (tr *TeamRepository) Create(team *entities.Team) error {
	ret := tr.Called(team)
	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Team) error); ok {
		r0 = rf(team)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}
	return r0
}

func (tr *TeamRepository) Update(team *entities.Team) error {
	ret := tr.Called(team)
	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Team) error); ok {
		r0 = rf(team)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}
	return r0
}

func (tr *TeamRepository) FindByID(id int) (*entities.Team, error) {
	return nil, nil
}

func (tr *TeamRepository) SearchTeams(keyword string) ([]*entities.Team, error) {
	return nil, nil
}
