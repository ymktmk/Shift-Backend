package domain

import "time"

type Users []User

type User struct {
    ID   uint    `gorm:"primaryKey,autoincrement" json:"id,omitempty"`
    UID  string  `gorm:"size:28;not null;unique" json:"uid,omitempty" validate:"required"`
    Name  string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
    Email string `gorm:"size:255;not null;unique" json:"email,omitempty" validate:"required,email"`
    CompanyID int `gorm:"not null" json:"company_id,omitempty"`
    Company Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
    CreatedAt *time.Time  `json:"-,omitempty"`
    UpdatedAt *time.Time  `json:"-,omitempty"`
	DeletedAt *time.Time  `json:"-,omitempty"`
}

type UserUpdateRequest struct {
    Name  string `json:"name" validate:"required"`
}

type UserUpdateResponse struct {
    UID  string   `gorm:"size:28;not null;unique" json:"uid"`
    Name  string `gorm:"size:255;not null" json:"name" validate:"required"`
    Email string `gorm:"size:255;not null;unique" json:"email" validate:"email"`
}