package domain

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name" required:"true"`
    Email string `json:"email"`
}

type Users []User