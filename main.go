package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ymktmk/Shift-Backend/infrastructure"
)

func main() {
	e := infrastructure.Routing()
	// middleware
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
        AllowMethods: []string{echo.GET,echo.POST},
    }))
	e.Logger.Fatal(e.Start(":9000"))
}