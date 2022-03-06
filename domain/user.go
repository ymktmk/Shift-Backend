package domain

import (
	"gorm.io/gorm"
)

type Users []User

type User struct {
    gorm.Model
    UID  string  `gorm:"size:28;not null;unique" json:"uid,omitempty" validate:"required"`
    Name  string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
    Email string `gorm:"size:255;not null;unique" json:"email,omitempty" validate:"required,email"`
    CompanyID int `gorm:"not null" json:"company_id,omitempty"`
    Company Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
}

type UserUpdateRequest struct {
    Name  string `json:"name" validate:"required"`
}

type UserUpdateResponse struct {
    UID  string   `json:"uid"`
    Name  string `json:"name"`
    Email string `json:"email"`
}