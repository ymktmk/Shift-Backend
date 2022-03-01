package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/ymktmk/Shift-Backend/infrastructure"
)

func main() {
	echo := infrastructure.Routing()
	echo.Use(middleware.CORS())
	echo.Logger.Fatal(echo.Start(":8080"))
}