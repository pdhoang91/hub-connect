package repo

import (
	"hub-connect/internal/entities"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepo{DB: db}
}

func (ur *UserRepo) Create(user *entities.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepo) Update(user *entities.User) error {
	// need handle update time
	//now := time.Now()
	return ur.DB.Model(&entities.User{}).Where("id = ?", user.ID).Updates(entities.User{
		TeamID: user.TeamID,
		//UpdatedAt: &now,
	}).Error
}

func (ur *UserRepo) FindByID(id int) (*entities.User, error) {
	var user entities.User
	if err := ur.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
