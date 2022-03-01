package domain

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name" required:"true"`
}

type Users []User