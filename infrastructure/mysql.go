package infrastructure

import (
	"github.com/ymktmk/Shift-Backend/interfaces/controllers"
	"github.com/labstack/echo"
)

func Routing() *echo.Echo {
	// user controller
	userController := controllers.NewUserController(NewSqlHandler())
	// routing
	echo := echo.New()
	echo.POST("/users/create", userController.CreateUser)
	echo.GET("/users/:id", userController.GetUser)
	echo.GET("/users",userController.GetAllUsers)
	return echo
}