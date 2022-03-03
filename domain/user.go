package domain

import (
    "time"
)

type User struct {
    ID   int  `gorm:"primary_key" json:"id"`
    UID  string   `gorm:"size:28;not null;unique" json:"uid"`
    Name  string `gorm:"size:255;not null" json:"name,omitempty"`
    Email string `gorm:"size:255;not null;unique" json:"email,omitempty"`
    // Password  string `gorm:"size:255;not null" json:"password,omitempty"`
    CreatedAt *time.Time  `json:"created_at"`
    UpdatedAt *time.Time  `json:"updated_at"`
    DeletedAt *time.Time  `json:"deleted_at"`
}

type Users []User