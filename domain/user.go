package domain

import "github.com/ymktmk/Shift-Backend/domain/gorm"

type User struct {
    gorm.Model
    UID  string   `gorm:"size:28;not null;unique" json:"uid"`
    Name  string `gorm:"size:255;not null" json:"name,omitempty"`
    Email string `gorm:"size:255;not null;unique" json:"email,omitempty"`
    // CompanyID int `gorm:"not null" json:"company_id"`
    // Company Company
    // Password  string `gorm:"size:255;not null" json:"password,omitempty"`
}

type Users []User