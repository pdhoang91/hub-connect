// Package usecase defines the business logic for managing users.
package usecase

import (
	"fmt"
	"hub-connect/internal/entities"
	repo "hub-connect/internal/repository"
	iError "hub-connect/pkg/error"
	"sync"
)

// IUserUseCase implements the UserUseCase interface and contains the business logic for user-related operations.
type IUserUseCase struct {
	UserRepository repo.UserRepository
	TeamRepository repo.TeamRepository
}

// NewUserUseCase creates a new instance of IUserUseCase with the provided UserRepository.
func NewUserUseCase(userRepo repo.UserRepository, teamRepo repo.TeamRepository) *IUserUseCase {
	return &IUserUseCase{
		UserRepository: userRepo,
		TeamRepository: teamRepo,
	}
}

// CreateUser handles the creation of a new user.
func (iuc *IUserUseCase) CreateUser(name string, email string) (*entities.User, error) {
	user := &entities.User{
		Name:  name,
		Email: email,
	}
	if err := iuc.UserRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates a user's associated team.
func (iuc *IUserUseCase) UpdateUser(id, teamID int) (*entities.User, error) {
	user := &entities.User{
		ID:     id,
		TeamID: &teamID,
	}
	err := iuc.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID retrieves user information by its ID.
func (iuc *IUserUseCase) GetUserByID(id int) (*entities.User, error) {
	user, err := iuc.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UserJoinTeam associates a user with a team.
func (iuc *IUserUseCase) UserJoinTeam(userID, teamID int) error {
	var wg sync.WaitGroup
	var userErr, teamErr error

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, userErr = iuc.UserRepository.FindByID(userID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, teamErr = iuc.TeamRepository.FindByID(teamID)
	}()

	wg.Wait()

	if userErr != nil {
		return iError.NewError(fmt.Sprintf("UserID [%d] not found, [%s]", userID, userErr))
	}

	if teamErr != nil {
		return iError.NewError(fmt.Sprintf("TeamID [%d] not found, [%s]", teamID, teamErr))
	}

	user := entities.User{
		ID:     userID,
		TeamID: &teamID,
	}

	err := iuc.UserRepository.Update(&user)
	if err != nil {
		return err
	}

	return nil
}
