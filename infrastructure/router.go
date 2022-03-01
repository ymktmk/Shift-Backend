package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
    "github.com/ymktmk/Shift-Backend/interfaces/controllers"
)

func Routing() *echo.Echo {
	// user controller
	userController := controllers.NewUserController(NewSqlHandler())
	// echo instance
    e := echo.New()
    // middleware
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
        AllowMethods: []string{echo.GET,echo.POST},
    }))
    e.Use(FirebaseAuth)
    // routing
    // -H 'Authorization: Bearer XXXXXXX'
	e.POST("/users/create", userController.CreateUser)
	e.GET("/users/:id", userController.GetUser)
    e.GET("/users",userController.GetAllUsers)
	return e
}
