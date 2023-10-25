package usecase

import "hub-connect/internal/entities"

// HubUseCase provides an interface for hub-related use cases.
type HubUseCase interface {
	CreateHub(name, location string) (*entities.Hub, error)
	GetHubByID(id int) (*entities.Hub, error)
	SearchHubs(keyword string) ([]*entities.Hub, error)
}

// TeamUseCase provides an interface for team-related use cases.
type TeamUseCase interface {
	CreateTeam(name, teamType string) (*entities.Team, error)
	UpdateTeam(id, hubID int) (*entities.Team, error)
	GetTeamByID(id int) (*entities.Team, error)
	TeamJoinHub(teamID, hubID int) error
	SearchTeams(keyword string) ([]*entities.Team, error)
}

// UserUseCase provides an interface for user-related use cases.
type UserUseCase interface {
	CreateUser(name string, email string) (*entities.User, error)
	UpdateUser(id, teamID int) (*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	UserJoinTeam(userID, teamID int) error
}
