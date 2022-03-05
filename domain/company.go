package domain

import "github.com/ymktmk/Shift-Backend/domain/gorm"

type Company struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
	Users []User `gorm:"foreignKey:CompanyID" json:"users,omitempty"`
}

type Companies []Company