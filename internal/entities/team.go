package entities

import "time"

type Team struct {
	ID        int        `json:"id" gorm:"column:id;"`
	Name      string     `json:"name" gorm:"column:name;"`
	Type      string     `json:"type" gorm:"column:type;"`
	HubID     *int       `json:"hub_id" gorm:"column:hub_id;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Users     []*User    `json:"users" gorm:"foreignkey:TeamID; references:ID"`
}

func (Team) TableName() string {
	return "team"
}
