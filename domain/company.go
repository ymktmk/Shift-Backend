package domain

import (
	"gorm.io/gorm"
)

type Companies []Company

type Company struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
	Users []User `gorm:"foreignKey:CompanyID" json:"users,omitempty"`
}

type CompanyRequest struct {
	ID   uint `json:"-"`
	Name string `json:"name" validate:"required"` 
}

// type CompanyResponse struct {
// 	ID   uint `gorm:"primaryKey,autoincrement" json:"id,omitempty"`
// 	Name string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"` 
// }