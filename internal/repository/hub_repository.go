package repo

import (
	"hub-connect/internal/entities"

	"gorm.io/gorm"
)

type HubRepo struct {
	DB *gorm.DB
}

func NewHubRepository(db *gorm.DB) HubRepository {
	return &HubRepo{DB: db}
}

func (hr *HubRepo) Create(hub *entities.Hub) error {
	return hr.DB.Create(hub).Error
}

func (hr *HubRepo) FindByID(id int) (*entities.Hub, error) {
	var hub entities.Hub
	if err := hr.DB.First(&hub, id).Error; err != nil {
		return nil, err
	}
	return &hub, nil
}

func (hr *HubRepo) SearchHubs(keyword string) ([]*entities.Hub, error) {
	var hubs []*entities.Hub
	if err := hr.DB.Where("name LIKE ?", "%"+keyword+"%").
		Preload("Teams").
		Find(&hubs).Error; err != nil {
		return nil, err
	}

	return hubs, nil
}
