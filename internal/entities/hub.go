package entities

import "time"

type Hub struct {
	ID        int        `json:"id" gorm:"column:id;"`
	Name      string     `json:"name"  gorm:"column:name;"`
	Location  string     `json:"location"  gorm:"column:location;"`
	CreatedAt *time.Time `json:"created_at"  gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at"  gorm:"column:updated_at;"`
	Teams     []*Team    `json:"teams" gorm:"foreignkey:HubID; references:ID"`
}

func (Hub) TableName() string {
	return "hub"
}
