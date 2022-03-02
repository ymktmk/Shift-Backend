package domain

type User struct {
    ID   int    `gorm:"primary_key" json:"id"`
    Name  string `json:"name"`
}

type Users []User