package domain

import (
	"time"

	"github.com/ymktmk/Shift-Backend/domain/gorm"
)

type Shift struct {
	gorm.Model
	Date *time.Time `json:"date,omitempty"`
	Vacation bool `json:"vacation,omitempty"`
	Require int `json:"require,omitempty"`
	UserID int `json:"user_id,omitempty"`
}

type ShiftCreateRequest struct {
    Date *time.Time `json:"date,omitempty"`
	Vacation bool `json:"vacation,omitempty"`
	Require int `json:"require,omitempty"`
}