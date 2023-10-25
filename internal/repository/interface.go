package repo

import "hub-connect/internal/entities"

type HubRepository interface {
	Create(hub *entities.Hub) error
	FindByID(id int) (*entities.Hub, error)
	SearchHubs(keyword string) ([]*entities.Hub, error)
}

type TeamRepository interface {
	Create(team *entities.Team) error
	Update(team *entities.Team) error
	FindByID(id int) (*entities.Team, error)
	SearchTeams(keyword string) ([]*entities.Team, error)
}

type UserRepository interface {
	Create(user *entities.User) error
	Update(user *entities.User) error
	FindByID(id int) (*entities.User, error)
}
