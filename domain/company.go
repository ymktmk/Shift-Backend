package domain

import "time"

type Companies []Company

type Company struct {
	ID   uint `gorm:"primaryKey,autoincrement" json:"id,omitempty"`
	Name string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
	Users []User `gorm:"foreignKey:CompanyID" json:"users,omitempty"`
	CreatedAt *time.Time  `json:"-,omitempty"`
    UpdatedAt *time.Time  `json:"-,omitempty"`
	DeletedAt *time.Time  `json:"-,omitempty"`
}

type CompanyRequest struct {
	ID   uint `json:"-"`
	Name string `json:"name" validate:"required"` 
}

// type CompanyResponse struct {
// 	ID   uint `gorm:"primaryKey,autoincrement" json:"id,omitempty"`
// 	Name string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"` 
// }