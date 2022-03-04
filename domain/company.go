package domain

import "github.com/ymktmk/Shift-Backend/domain/gorm"

type Company struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name,omitempty"`
	// Users []User `gorm:"foreignKey:CompanyID"`
}
