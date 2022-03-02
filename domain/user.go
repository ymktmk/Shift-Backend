package domain

type User struct {
    ID   int    `json:"id"`
    Name  string `json:"name" validate:"required, min=5"`
}

type Users []User