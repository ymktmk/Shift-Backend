package main

import (
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/ymktmk/Shift-Backend/domain"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
    user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	protcol := "tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")"
	name := os.Getenv("MYSQL_DATABASE") + "?charset=utf8&parseTime=true&loc=Local"
	connect := user + ":" + password + "@" + protcol + "/" + name
    conn, err := gorm.Open("mysql", connect)
	if err != nil {
        panic(err)
    }
    defer func() {
        if err := conn.Close(); err != nil {
            panic(err)
        }
    }()
    conn.LogMode(true)
    if err := conn.DB().Ping(); err != nil {
        panic(err)
	}
	
	// migration
    conn.AutoMigrate(
        &domain.User{},
    )
}