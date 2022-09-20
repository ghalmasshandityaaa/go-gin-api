package entities

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-" gorm:"column:password"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at"`
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"index,column:delete_at"`
}
