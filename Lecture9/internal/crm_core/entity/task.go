package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Description      string    `gorm:"varchar(255);not null" json:"description"`
	DueDate          time.Time `json:"due_date"`
	AssignedTo       uuid.UUID `json:"assigned_to"`
	AssociatedDealID uint      `json:"associated_deal_id"`
}
