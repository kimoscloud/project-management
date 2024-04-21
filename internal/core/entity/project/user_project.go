package project

import (
	"gorm.io/gorm"
	"time"
)

type UserProject struct {
	ID        string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TeamID    string         `gorm:"type:uuid;primaryKey"`
	ProjectID string         `gorm:"type:uuid;primaryKey"`
	Project   Project        `gorm:"foreignKey:ProjectID;references:ID"`
	Status    string         `gorm:"type:varchar(255);not null"` // Status of the user in the project (pending, accepted, rejected)
	IsActive  bool           `gorm:"type:boolean;default:true"`  // If the team has access or not
	Role      string         `gorm:"type:varchar(255);not null"` // ProjectRole of the team in the project
	CreatedAt time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (UserProject) TableName() string {
	return "User_Projects"
}
