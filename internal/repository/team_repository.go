package repo

import (
	"hub-connect/internal/entities"

	"gorm.io/gorm"
)

type TeamRepo struct {
	DB *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &TeamRepo{DB: db}
}

func (tr *TeamRepo) Create(team *entities.Team) error {
	return tr.DB.Create(team).Error
}

func (tr *TeamRepo) Update(team *entities.Team) error {
	return tr.DB.Model(&entities.Team{}).Where("id = ?", team.ID).Updates(team).Error
}

func (tr *TeamRepo) FindByID(id int) (*entities.Team, error) {
	var team entities.Team
	if err := tr.DB.First(&team, id).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (tr *TeamRepo) SearchTeams(keyword string) ([]*entities.Team, error) {
	var teams []*entities.Team
	if err := tr.DB.Where("name LIKE ?", "%"+keyword+"%").
		Preload("Users").
		Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}
