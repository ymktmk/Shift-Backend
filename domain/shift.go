package domain

import (
	"time"

	"github.com/ymktmk/Shift-Backend/domain/gorm"
)

type Shift struct {
	gorm.Model
	Date *time.Time `json:"date,omitempty"`
	Vacation bool `gorm:"not null" json:"vacation,omitempty"`
	Require int `gorm:"size:28;not null" json:"require,omitempty"`
	UserID int `gorm:"not null" json:"user_id,omitempty"`
}