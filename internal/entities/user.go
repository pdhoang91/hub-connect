package entities

import "time"

type User struct {
	ID        int        `json:"id" gorm:"column:id;"`
	Name      string     `json:"name" gorm:"column:name;"`
	Email     string     `json:"email" gorm:"column:email;"`
	TeamID    *int       `json:"team_id" gorm:"column:team_id;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (User) TableName() string {
	return "user_info"
}
