package domain

import "github.com/ymktmk/Shift-Backend/domain/gorm"

// フィールドを完全に無視 `json:"-"`
type User struct {
    gorm.Model
    UID  string   `gorm:"size:28;not null;unique" json:"uid,omitempty" validate:"required"`
    Name  string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
    Email string `gorm:"size:255;not null;unique" json:"email,omitempty" validate:"required,email"`
    CompanyID int `gorm:"not null" json:"company_id,omitempty"`
    Company Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
}

type Users []User