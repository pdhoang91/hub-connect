package mocks

import (
	"hub-connect/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

type TeamUseCase struct {
	mock.Mock
}

// CreateTeam handles the creation of a new team.
func (_m *TeamUseCase) CreateTeam(name, teamType string) (*entities.Team, error) {

	ret := _m.Called(name, teamType)

	var r0 *entities.Team
	if rf, ok := ret.Get(0).(func(string, string) *entities.Team); ok {
		r0 = rf(name, teamType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, teamType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTeam updates a team's associated hub.
func (_m *TeamUseCase) UpdateTeam(id, hubID int) (*entities.Team, error) {
	// Need time to do
	// ToDo
	return nil, nil
}

// GetTeamByID retrieves team information by its ID.
func (_m *TeamUseCase) GetTeamByID(id int) (*entities.Team, error) {
	// Need time to do
	// ToDo
	return nil, nil
}

// TeamJoinHub associates a team with a hub.
func (_m *TeamUseCase) TeamJoinHub(teamID, hubID int) error {
	// Need time to do
	// ToDo
	return nil
}

// SearchTeams searches for teams based on a keyword.
func (_m *TeamUseCase) SearchTeams(keyword string) ([]*entities.Team, error) {
	// Need time to do
	// ToDo
	return nil, nil
}
