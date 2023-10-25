// Package usecase defines the business logic for managing teams.
package usecase

import (
	"fmt"
	"hub-connect/internal/entities"
	repo "hub-connect/internal/repository"
	iError "hub-connect/pkg/error"
	"sync"
)

// ITeamUseCase implements the TeamUseCase interface and contains the business logic for team-related operations.
type ITeamUseCase struct {
	TeamRepository repo.TeamRepository
	HubRepository  repo.HubRepository
}

// NewTeamUseCase creates a new instance of ITeamUseCase with the provided TeamRepository.
func NewTeamUseCase(teamRepo repo.TeamRepository, hubRepo repo.HubRepository) *ITeamUseCase {
	return &ITeamUseCase{
		TeamRepository: teamRepo,
		HubRepository:  hubRepo,
	}
}

// CreateTeam handles the creation of a new team.
func (ti *ITeamUseCase) CreateTeam(name, teamType string) (*entities.Team, error) {
	team := &entities.Team{
		Name: name,
		Type: teamType,
	}
	if err := ti.TeamRepository.Create(team); err != nil {
		return nil, err
	}
	return team, nil
}

// UpdateTeam updates a team's associated hub.
func (ti *ITeamUseCase) UpdateTeam(id, hubID int) (*entities.Team, error) {
	team := &entities.Team{
		ID:    id,
		HubID: &hubID,
	}
	if err := ti.TeamRepository.Update(team); err != nil {
		return nil, err
	}
	return team, nil
}

// GetTeamByID retrieves team information by its ID.
func (ti *ITeamUseCase) GetTeamByID(id int) (*entities.Team, error) {
	team, err := ti.TeamRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return team, nil
}

// TeamJoinHub associates a team with a hub.
func (ti *ITeamUseCase) TeamJoinHub(teamID, hubID int) error {

	var teamErr, hubErr error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, teamErr = ti.TeamRepository.FindByID(teamID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, hubErr = ti.HubRepository.FindByID(hubID)
	}()

	wg.Wait()

	if teamErr != nil {
		return iError.NewError(fmt.Sprintf("TeamID [%d] not found, [%s]", teamID, teamErr))
	}

	if hubErr != nil {
		return iError.NewError(fmt.Sprintf("HubID [%d] not found, [%s]", hubID, hubErr))
	}

	team := &entities.Team{
		ID:    teamID,
		HubID: &hubID,
	}

	if err := ti.TeamRepository.Update(team); err != nil {
		return err
	}

	return nil
}

// SearchTeams searches for teams based on a keyword.
func (ti *ITeamUseCase) SearchTeams(keyword string) ([]*entities.Team, error) {
	data, err := ti.TeamRepository.SearchTeams(keyword)
	if err != nil {
		return nil, err
	}
	return data, nil
}