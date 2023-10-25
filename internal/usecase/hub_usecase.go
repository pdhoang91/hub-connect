// Package usecase defines the business logic for managing hubs.
package usecase

import (
	"hub-connect/internal/entities"
	repo "hub-connect/internal/repository"
)

// IHubUseCase implements the HubUseCase interface and contains the business logic for hub-related operations.
type IHubUseCase struct {
	HubRepository repo.HubRepository
}

// NewHubUseCase creates a new instance of IHubUseCase with the provided HubRepository.
func NewHubUseCase(hubRepo repo.HubRepository) *IHubUseCase {
	return &IHubUseCase{HubRepository: hubRepo}
}

// CreateHub handles the creation of a new hub.
func (hi *IHubUseCase) CreateHub(name, location string) (*entities.Hub, error) {
	hub := &entities.Hub{
		Name:     name,
		Location: location,
	}
	err := hi.HubRepository.Create(hub)
	if err != nil {
		return nil, err
	}
	return hub, nil
}

// GetHubByID retrieves hub information by its ID.
func (hi *IHubUseCase) GetHubByID(id int) (*entities.Hub, error) {
	hub, err := hi.HubRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return hub, nil
}

// SearchHubs searches for hubs based on a keyword.
func (hi *IHubUseCase) SearchHubs(keyword string) ([]*entities.Hub, error) {
	data, err := hi.HubRepository.SearchHubs(keyword)
	if err != nil {
		return nil, err
	}
	return data, nil
}
